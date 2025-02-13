package postgres

import (
	"database/sql"
	"fmt"

	"github.com/quanergyO/avito_assingment/types"
)

type Auth struct {
	db *sql.DB
}

func NewAuth(db *sql.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUser(user types.UserType) (int, error) {

	return 0, fmt.Errorf("repository CreateUser not implemented")
}

func (r *Auth) GetUser(username, password string) (types.UserType, error) {

	return types.UserType{}, fmt.Errorf("repository GetUser not implemented")
}
