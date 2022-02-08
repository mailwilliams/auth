package models

type User struct {
	UserID        uint64 `json:"user_id"`
	WalletAddress string `json:"wallet_address"`
	Password      string `json:"password"`
	PasswordMatch string `json:"password_match"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Mobile        string `json:"mobile"`
}
