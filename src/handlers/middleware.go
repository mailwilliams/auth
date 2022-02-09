package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

//	IsAuthenticated is the middleware method that will be used to separate functions into
//	categories that are permitted if the user has a valid jwt token
func (handler *Handler) IsAuthenticated(c *fiber.Ctx) error {
	cookie := handler.GetCookie(c)
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return handler.ErrResponse(c, fiber.StatusUnauthorized, fiber.Map{
			"message": "user unauthenticated",
		})
	}

	return c.Next()
}

//	GetUserID is a middleware method that will return a user's ID
//	by parsing the token stored in the cookie
func (handler *Handler) GetUserID(c *fiber.Ctx) (uint64, error) {
	token, err := jwt.ParseWithClaims(handler.GetCookie(c), &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return 0, err
	}

	userID, err := strconv.ParseInt(token.Claims.(*jwt.StandardClaims).Subject, 10, 64)

	return uint64(userID), nil
}
