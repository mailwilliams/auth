package handlers

import (
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
	POST	/api/register

Description:
	-	Receives user information from the payload
	-	Checks to see that password and password_confirm are the same
	=	Creates user in the database
	-	Generate JWT token
	-	Save token as cookie

Payload:
	{
		"email": "liam@sundial.ai1",
		"password": "abc123",
		"password_confirm": "abc123",
		"first_name": "Liam",
		"last_name": "Williams"
	}

Response:
	-	200 OK
		-	User successfully logged in, cookie set

	-	400 Bad Request
		-	password does not match password_match

	-	401 Not Found
		-	User not found
		-	Invalid credentials

	-	500 Internal Server Error
		-	Database error
		-	JWT token generation error
*/

func (handler *Handler) Register(c *fiber.Ctx) error {
	var requestBody map[string]string

	//	parsing payload and storing it in map[string]string
	if err := c.BodyParser(&requestBody); err != nil {
		return err
	}

	//	assigning values from payload to a new user struct
	user := models.User{
		WalletAddress:   requestBody["wallet_address"],
		Password:        []byte(requestBody["password"]),
		PasswordConfirm: []byte(requestBody["password_confirm"]),
		FirstName:       requestBody["first_name"],
		LastName:        requestBody["last_name"],
		Email:           requestBody["email"],
		Mobile:          requestBody["mobile"],
	}

	//	comparing to confirm the user is entering the correct password
	if string(user.Password) != string(user.PasswordConfirm) {
		return handler.ErrResponse(c, http.StatusBadRequest, fiber.Map{
			"message": "passwords do not match",
		})
	}

	//	generating a hash of the confirmed password
	if err := user.SetPassword(user.Password); err != nil {
		return handler.ErrResponse(c, http.StatusBadRequest, fiber.Map{
			"message": err.Error(),
		})
	}

	//	inserting the new user into the database
	result, err := handler.DB.ExecContext(handler.ctx, SQL.Register,
		user.WalletAddress,
		user.Password,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Mobile,
	)
	if err != nil {
		return handler.ErrResponse(c, http.StatusBadRequest, fiber.Map{
			"message": err.Error(),
		})
	}

	//	returning the newly created userID to store in the signed jwt token
	userID, err := result.LastInsertId()
	if err != nil {
		return handler.ErrResponse(c, http.StatusBadRequest, fiber.Map{
			"message": err.Error(),
		})
	}

	//	creating a new signed jwt token
	signedString, err := handler.GenerateJWT(models.JWT{
		StandardClaims: jwt.StandardClaims{
			Subject:   strconv.Itoa(int(userID)),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
		WalletAddress: user.WalletAddress,
	})
	if err != nil {
		return handler.ErrResponse(c, http.StatusBadRequest, fiber.Map{
			"message": "Invalid credentials",
		})
	}

	//	setting the user's "jwt" cookie to the signed jwt token string
	handler.SetCookie(c, signedString, time.Now().Add(time.Hour*24))

	return handler.SuccessResponse(c, fiber.StatusCreated, nil)
}
