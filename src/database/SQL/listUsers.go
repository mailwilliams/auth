package SQL

const ListUsers = `
SELECT
	user_id,
	wallet_address,
	first_name,
	last_name
FROM
	auth.users;`
