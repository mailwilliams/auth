package handlers

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/mailwilliams/auth/src/models"
	"net/http"
	"strconv"
	"time"
)

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

	if err := handler.DB.QueryRowContext(handler.ctx, loginSQL(), user.Email).Scan(&user.UserID, &user.Password); err != nil {
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
		return handler.ErrResponse(c, fiber.StatusBadRequest, fiber.Map{
			"message": "Invalid credentials",
		})
	}

	signedString, err := handler.GenerateJWT(struct {
		jwt.StandardClaims
		//	add more here
		//	EX:
		//	Scope string
	}{
		StandardClaims: jwt.StandardClaims{
			Subject:   strconv.Itoa(int(user.UserID)),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			// Scope: "admin"
		},
	})
	if err != nil {
		return handler.ErrResponse(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	handler.SetCookie(c, signedString)

	return handler.SuccessResponse(c, http.StatusOK, nil)
}

func loginSQL() string {
	return `
SELECT
	user_id,
	password
FROM
	auth.users
WHERE
	email = ?;`
}
