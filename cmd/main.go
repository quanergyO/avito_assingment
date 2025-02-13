package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/quanergyO/avito_assingment/internal/handler"
	"github.com/quanergyO/avito_assingment/internal/repository"
	"github.com/quanergyO/avito_assingment/internal/repository/postgres"
	"github.com/quanergyO/avito_assingment/internal/service"
	"github.com/quanergyO/avito_assingment/server"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		slog.Error("Error: init configs", slog.Any("error", err))
		os.Exit(1)
	}

	if err := godotenv.Load(); err != nil {
		slog.Error("Error: loading env variables", slog.Any("error", err))
		os.Exit(1)
	}

	db, err := postgres.NewDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		slog.Error("Error: failed to init db connection", slog.Any("error", err))
		os.Exit(1)
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	serv := new(server.Server)
	go func() {
		if err := serv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			slog.Error("Error: failed to start server on port:", viper.GetString("db.host"), err.Error())
			os.Exit(1)
		}
	}()

	slog.Info("Start server")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := serv.ShutDown(context.Background()); err != nil {
		slog.Error("Error occured on server shutting down", slog.Any("error", err))
		os.Exit(1)
	}
	slog.Info("Server shutting down")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
