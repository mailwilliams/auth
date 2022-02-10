package SQL

const Register = `
INSERT INTO auth.users
	(created_at, wallet_address, password, first_name, last_name, email, mobile)
VALUES
	(NOW(), ?, ?, ?, ?, ?, ?);`
