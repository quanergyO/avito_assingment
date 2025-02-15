package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/quanergyO/avito_assingment/types"
)

type Auth struct {
	db *sql.DB
}

func NewAuth(db *sql.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUser(user types.UserType) (int, error) {
	const op = "postgres.Auth.CreateUser"

	log := slog.With(
		slog.String("op", op),
	)
	log.Info("Call CreateUser")

	query := `
        INSERT INTO users (username, password_hash)
        VALUES ($1, $2)
        RETURNING id
    `

	var userId int
	err := r.db.QueryRow(query, user.Username, user.Password).Scan(&userId)

	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (r *Auth) GetUser(username, password string) (types.UserType, error) {
	const op = "postgres.Auth.GetUser"

	log := slog.With(
		slog.String("op", op),
	)
	log.Info("Call GetUser")

	var user types.UserType
	query := fmt.Sprintf("SELECT id, username, password_hash, coins FROM %s WHERE username=$1 AND password_hash=$2", userTable)
	err := r.db.QueryRow(query, username, password).Scan(&user.Id, &user.Username, &user.Password, &user.Coins)
	if err != nil {
		return types.UserType{}, err
	}

	return user, nil
}
