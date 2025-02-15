package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const (
	userTable = "users"
)

func NewDB(cfg Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName)
	slog.Info(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		slog.Error("error open connection with db", slog.Any("error", err))
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		slog.Error("error ping to db", slog.Any("error", err))
		return nil, err
	}

	err = makeMigrations(cfg)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func makeMigrations(cfg Config) error {
	wd, err := os.Getwd()
	if err != nil {
		slog.Error("Can't get current working directory", slog.Any("error", err))
		return err
	}
	migrationPath := filepath.Join(wd, "schema")
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)

	m, err := migrate.New(
		"file://"+migrationPath,
		databaseURL,
	)
	if err != nil {
		slog.Error("failed to connect db", slog.Any("error", err))
		return err
	}

	err = m.Up()
	if err != nil {
		slog.Error("Migration is not used", err.Error())
	}

	return nil
}
