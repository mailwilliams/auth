package SQL

const UpdateUserInfo = `
UPDATE
	auth.users
SET
	updated_at = NOW(),
	email = ?,
	first_name = ?,
	last_name = ?,
	mobile = ?
WHERE
	user_id = ?;`
