package models

type User struct {
	UserID        uint64 `json:"user_id"`
	WalletAddress string `json:"wallet_address"`
	Password      []byte `json:"password"`
	PasswordMatch []byte `json:"password_match" gorm:"-"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Mobile        string `json:"mobile"`
}
