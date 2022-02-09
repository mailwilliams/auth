package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/mailwilliams/auth/src/models"
	"net/http"
	"strconv"
	"time"
)

func (handler *Handler) Register(c *fiber.Ctx) error {
	var requestBody map[string]string

	if err := c.BodyParser(&requestBody); err != nil {
		return err
	}

	user := models.User{
		WalletAddress: requestBody["wallet_address"],
		Password:      []byte(requestBody["password"]),
		PasswordMatch: []byte(requestBody["password_match"]),
		FirstName:     requestBody["first_name"],
		LastName:      requestBody["last_name"],
		Email:         requestBody["email"],
		Mobile:        requestBody["mobile"],
	}

	if string(user.Password) != string(user.PasswordMatch) {
		return handler.ErrResponse(c, http.StatusBadRequest, fiber.Map{
			"message": "passwords do not match",
		})
	}

	if err := user.SetPassword(user.Password); err != nil {
		return handler.ErrResponse(c, http.StatusBadRequest, fiber.Map{
			"message": err.Error(),
		})
	}

	result, err := handler.DB.ExecContext(handler.ctx, registerSQL(),
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

	userID, err := result.LastInsertId()
	if err != nil {
		return handler.ErrResponse(c, http.StatusBadRequest, fiber.Map{
			"message": err.Error(),
		})
	}

	signedString, err := handler.GenerateJWT(struct {
		jwt.StandardClaims
		//	add more here
		//	EX:
		//	Scope string
	}{
		StandardClaims: jwt.StandardClaims{
			Subject:   strconv.Itoa(int(userID)),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			// Scope: "admin"
		},
	})
	if err != nil {
		return handler.ErrResponse(c, http.StatusBadRequest, fiber.Map{
			"message": "Invalid credentials",
		})
	}

	handler.SetCookie(c, signedString)

	return handler.SuccessResponse(c, fiber.StatusCreated, nil)
}

func registerSQL() string {
	return `
INSERT INTO auth.users
	(wallet_address, password, first_name, last_name, email, mobile)
VALUES
	(?, ?, ?, ?, ?, ?);`
}
