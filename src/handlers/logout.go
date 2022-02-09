package handlers

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

/*
Method:
	DELETE	/api/logout

Description:
	-	Creates a new cookie that has no value and is already expired
	-	Assigns expired cookie to user to clear existing cookie

Payload:
	-	N/A

Response:
	-	204 No Content
		-	User successfully logged out, cookie set to empty and expired
*/

func (handler *Handler) Logout(c *fiber.Ctx) error {
	handler.SetCookie(c, "", time.Now().Add(-time.Hour))

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
