package handlers

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/mailwilliams/auth/src/database/SQL"
	"github.com/mailwilliams/auth/src/models"
	"net/http"
	"strconv"
	"time"
)

/*
Method:
	POST	/api/login

Description:
	-	Receives an email and a password from the payload
	-	Compares it against any records in the database
	-	Generate JWT token
	-	Save token as cookie

Payload:
	{
		"email": "liam@sundial.ai",
		"password": "abc123"
	}

Response:
	-	200 OK
		-	User successfully logged in, cookie set

	-	401 Not Found
		-	User not found
		-	Invalid credentials

	-	500 Internal Server Error
		-	Database error
		-	JWT token generation error
*/

func (handler *Handler) Login(c *fiber.Ctx) error {
	var requestBody map[string]string

	if err := c.BodyParser(&requestBody); err != nil {
		return handler.ErrResponse(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	user := models.User{
		Email: requestBody["email"],
	}

	if err := handler.DB.QueryRowContext(handler.ctx, SQL.Login, user.Email).Scan(&user.UserID, &user.Password); err != nil {
		switch err {
		case sql.ErrNoRows:
			return handler.ErrResponse(c, fiber.StatusNotFound, fiber.Map{
				"message": "user not found",
			})
		default:
			return handler.ErrResponse(c, fiber.StatusInternalServerError, fiber.Map{
				"message": err.Error(),
			})
		}
	}

	if err := user.ComparePassword(requestBody["password"]); err != nil {
		return handler.ErrResponse(c, fiber.StatusNotFound, fiber.Map{
			"message": "user not found",
		})
	}

	signedString, err := handler.GenerateJWT(models.JWT{
		StandardClaims: jwt.StandardClaims{
			Subject:   strconv.Itoa(int(user.UserID)),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
		WalletAddress: user.WalletAddress,
	})
	if err != nil {
		return handler.ErrResponse(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	handler.SetCookie(c, signedString, time.Now().Add(time.Hour*24))

	return handler.SuccessResponse(c, http.StatusOK, nil)
}
