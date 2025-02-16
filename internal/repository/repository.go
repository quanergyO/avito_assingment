package repository

import (
	"database/sql"

	"github.com/quanergyO/avito_assingment/internal/repository/postgres"
	"github.com/quanergyO/avito_assingment/types"
)

type Authorization interface {
	CreateUser(user types.SignInInput) (int, error)
	GetUser(username, password string) (types.UserDAO, error)
}

type User interface {
	GetUserInfo(userID int) (types.UserInfo, error)
	SendCoins(senderID, receiverID int, amount int) error
	BuyItem(userID int, itemName string) error
}

type Repository struct {
	Authorization
	User
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuth(db),
		User:          postgres.NewUserRepository(db),
	}
}
