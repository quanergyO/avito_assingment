package types

import "github.com/dgrijalva/jwt-go"

type UserType struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
