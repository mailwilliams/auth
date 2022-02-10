package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mailwilliams/auth/src/database/SQL"
	"github.com/mailwilliams/auth/src/models"
)

/*
Method:
	PUT	/api/users/me/wallet

Description:
	-	Receives user's wallet address from payload
	-	Finds the current user's ID by parsing the jwt token in the cookie
	-	Runs a query to update the user's wallet_address in the database
	-	Return OK

Payload:
	{
		"wallet_address": "0x614C1dd6FdEB6Ae19B4C05c4a39C6296FA885215",
	}

Response:
	-	200 OK
		-	User information updated successfully, returning updated user JSON

	-	401 Unauthorized
		-	Couldn't get userID from cookie

	-	500 Internal Server Error
		-	Database error
		-	Couldn't parse body
*/

func (handler *Handler) UpdateWalletAddress(c *fiber.Ctx) error {
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
		UserID:        userID,
		WalletAddress: requestBody["wallet_address"],
	}

	//	updating the user with new password
	_, err = handler.DB.ExecContext(handler.ctx, SQL.UpdateWalletAddress, user.WalletAddress, user.UserID)
	if err != nil {
		return handler.ErrResponse(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	return handler.SuccessResponse(c, fiber.StatusOK, nil)
}
