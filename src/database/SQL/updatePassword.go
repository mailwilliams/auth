package SQL

const UpdatePassword = `
UPDATE
	auth.users
SET 
	password = ?
WHERE
	user_id = ?;`
