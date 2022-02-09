package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

//	User is the struct type of the database table auth.users
//	It is also used in the handler methods to return JSON payloads
type User struct {
	UserID          uint64         `json:"user_id" gorm:"primaryKey:user_id"`
	CreatedAt       time.Time      `json:"-"`
	UpdatedAt       time.Time      `json:"-"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`
	WalletAddress   string         `json:"wallet_address" gorm:"unique size:1024"`
	Password        []byte         `json:"password,omitempty" form:"size:1024"`
	PasswordConfirm []byte         `json:"password_confirm,omitempty" gorm:"-"`
	FirstName       string         `json:"first_name"`
	LastName        string         `json:"last_name"`
	Email           string         `json:"email,omitempty" gorm:"unique size:256"`
	Mobile          string         `json:"mobile,omitempty" gorm:"unique size:256"`
}

func (user *User) SetPassword(password []byte) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, 12)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return nil
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
