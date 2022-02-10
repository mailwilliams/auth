package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mailwilliams/auth/src/database/SQL"
	"github.com/mailwilliams/auth/src/models"
)

/*
Method:
	PUT	/api/users/me

Description:
	-	Receives user information from the payload
	-	Finds the current user's ID by parsing the jwt token in the cookie
	-	Runs a query to update the user's information in the database
	-	Return updated user information

Payload:
	{
		"email": "liam@sundial.ai2",
		"password": "abc123",
		"password_match": "abc123",
		"first_name": "Liam",
		"last_name": "Williams"
	}

Response:
	-	200 OK
		-	User information updated successfully, returning updated user JSON

	-	401 Unauthorized
		-	Couldn't get userID from cookie

	-	500 Internal Server Error
		-	Database error
		-	Couldn't parse body

	{
		"email": "liam@sundial.ai2",
		"first_name": "Liam",
		"last_name": "Williams"
		"mobile": "+19254576277"
	}
*/

func (handler *Handler) UpdateInfo(c *fiber.Ctx) error {
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
		UserID:    userID,
		FirstName: requestBody["first_name"],
		LastName:  requestBody["last_name"],
		Email:     requestBody["email"],
		Mobile:    requestBody["mobile"],
	}

	//	updating the user with new information
	_, err = handler.DB.ExecContext(handler.ctx, SQL.UpdateUserInfo,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Mobile,
		user.UserID,
	)
	if err != nil {
		return handler.ErrResponse(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	return handler.SuccessResponse(c, fiber.StatusOK, user)
}
