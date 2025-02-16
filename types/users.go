package types

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserInfo struct {
	UserInfo        UserResponse
	PurchasesInfo   []PurchasesResponse
	TransactionInfo []TransactionsResponse
}

type SendCoinRequest struct {
	ReceiverId int `json:"receiver_id"`
	Amount     int `json:"amount"`
}

type UserResponse struct {
	Username string `json:"username" db:"username"`
	Coins    int    `json:"coins" db:"coins"`
}

type PurchasesResponse struct {
	ItemName string `json:"item_name"`
	Quantity int    `json:"quantity" db:"quantity"`
}

type TransactionsResponse struct {
	FromUserId int       `json:"from_user_id" db:"from_user_id"`
	ToUserId   int       `json:"to_user_id" db:"to_user_id"`
	Timestamp  time.Time `json:"timestamp" db:"timestamp"`
	Amount     int       `json:"amount" db:"amount"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
