package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

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
