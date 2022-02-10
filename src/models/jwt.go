package models

import "github.com/dgrijalva/jwt-go"

type JWT struct {
	jwt.StandardClaims
	//	add custom values here
	WalletAddress string
}
