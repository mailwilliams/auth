package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mailwilliams/auth/src/database/SQL"
	"github.com/mailwilliams/auth/src/models"
	"net/http"
)

/*
Method:
	PUT	/api/users/me/password

Description:
	-	Receives user's updated password from the payload
	-	Finds the current user's ID by parsing the jwt token in the cookie
	-	Runs a query to update the user's information in the database
	-	Return OK

Payload:
	{
		"password": "abc1234",
		"password_confirm": "abc1234",
	}

Response:
	-	200 OK
		-	User information updated successfully, returning OK

	-	401 Unauthorized
		-	Couldn't get userID from cookie

	-	500 Internal Server Error
		-	Database error
		-	Couldn't parse body
*/

func (handler *Handler) UpdatePassword(c *fiber.Ctx) error {
	var requestBody map[string]string

	//	saving payload as map[string]string
	if err := c.BodyParser(&requestBody); err != nil {
		return handler.ErrResponse(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	//	retrieving userID from cookie
	userID, err := handler.GetUserID(c)
	if err != nil {
		return handler.ErrResponse(c, fiber.StatusUnauthorized, fiber.Map{
			"message": err.Error(),
		})
	}

	//	assigning values to new user struct
	user := models.User{
		UserID:          userID,
		Password:        []byte(requestBody["password"]),
		PasswordConfirm: []byte(requestBody["password_confirm"]),
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

	//	updating the user with new password
	_, err = handler.DB.ExecContext(handler.ctx, SQL.UpdatePassword, user.Password, user.UserID)
	if err != nil {
		return handler.ErrResponse(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	return handler.SuccessResponse(c, fiber.StatusOK, nil)
}
