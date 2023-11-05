package entities

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	jwt.RegisteredClaims
	Phone  string `json:"phone"`
	UserId int64  `json:"user_id"`
}

type RegisteredUser struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}
