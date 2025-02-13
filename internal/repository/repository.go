package repository

import (
	"database/sql"

	"github.com/quanergyO/avito_assingment/types"
)

type Authorization interface {
	CreateUser(user types.UserType) (int, error)
	GetUser(username, password string) (types.UserType, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}
