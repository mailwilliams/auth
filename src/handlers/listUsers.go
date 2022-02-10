package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mailwilliams/auth/src/database/SQL"
	"github.com/mailwilliams/auth/src/models"
	"net/http"
)

/*
Method:
	GET	/api/users

Description:
	-	Retrieves all existing users

Payload:
	-	N/A

Response:
	-	200 OK
		-	All users returned in payload

	-	401 Unauthorized
		-	User not authenticated

	-	500 Internal Server Error
		-	Database error
		-	Data parsing error
*/

func (handler *Handler) ListUsers(c *fiber.Ctx) error {
	var users []models.User

	//	finding all users
	//	TODO: Add filtering
	rows, err := handler.DB.QueryContext(handler.ctx, SQL.ListUsers)
	if err != nil {
		return handler.ErrResponse(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	//	looping through all rows returned from SQL query
	for rows.Next() {
		var user models.User

		//	assigning the value of the current row to a new user struct
		if err := rows.Scan(&user.UserID, &user.WalletAddress, &user.FirstName, &user.LastName); err != nil {
			return handler.ErrResponse(c, fiber.StatusInternalServerError, fiber.Map{
				"message": err.Error(),
			})
		}

		//	adding that user to the list of users
		users = append(users, user)
	}

	return handler.SuccessResponse(c, http.StatusOK, users)
}
