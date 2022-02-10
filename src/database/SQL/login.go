package SQL

const Login = `
SELECT
	user_id,
	password
FROM
	auth.users
WHERE
	email = ?;`
