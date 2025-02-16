package types

import "time"

type UserDAO struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Coins    int    `json:"coins"`
}

type ItemDAO struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Price int    `json:"price" db:"price"`
}

type PurchasesDAO struct {
	Id        int       `json:"id" db:"id"`
	UserId    int       `json:"user_id" db:"user_id"`
	ItemId    int       `json:"item_id" db:"item_id"`
	Quantity  int       `json:"quantity" db:"quantity"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
}

type TransactionsDAO struct {
	Id         int       `json:"id" db:"id"`
	FromUserId int       `json:"from_user_id" db:"from_user_id"`
	ToUserId   int       `json:"to_user_id" db:"to_user_id"`
	Timestamp  time.Time `json:"timestamp" db:"timestamp"`
	Amount     int       `json:"amount" db:"amount"`
}
