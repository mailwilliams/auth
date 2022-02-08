package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mailwilliams/auth/src/models"
)

func (handler *Handler) Register(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	if user.Password != user.PasswordMatch {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	_, err := handler.DB.ExecContext(handler.ctx, `INSERT INTO users ()`)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(user)
}
