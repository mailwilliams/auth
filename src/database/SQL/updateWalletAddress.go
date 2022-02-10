package SQL

const UpdateWalletAddress = `
UPDATE
	auth.users
SET
	wallet_address = ?
WHERE
	user_id = ?;`
