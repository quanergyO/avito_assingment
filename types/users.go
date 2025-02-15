package types

import "github.com/dgrijalva/jwt-go"

type UserType struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Coins    int    `json:"coins"`
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserInfo struct {
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
