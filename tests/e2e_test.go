package tests

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/quanergyO/avito_assingment/internal/handler"
	"github.com/quanergyO/avito_assingment/internal/repository"
	"github.com/quanergyO/avito_assingment/internal/repository/postgres"
	"github.com/quanergyO/avito_assingment/internal/service"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func TestHandler_GetInfo(t *testing.T) {
	if err := initConfig(); err != nil {
		slog.Error("Error: init configs")
		os.Exit(1)
	}

	db, err := postgres.NewDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: "qwerty",
	})
	if err != nil {
		slog.Error("Error: failed to init db connection ")
		os.Exit(1)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	testTable := []struct {
		name                     string
		signInBody               string
		requestBody              string
		expectedStatusCodeVerify int
		expectedStatusCode       int
		expectedResponseBody     string
	}{
		{
			name:                     "GetInfo Ok",
			signInBody:               `{"username": "admin", "password": "admin"}`,
			requestBody:              ``,
			expectedStatusCodeVerify: 200,
			expectedStatusCode:       200,
			expectedResponseBody:     `{"data":{"UserInfo":{"username":"admin","coins":1000},"PurchasesInfo":[],"TransactionInfo":[]}}`,
		},
		{
			name:                     "Bad password token",
			signInBody:               `{"username": "admin", "password": "admin1"}`,
			requestBody:              ``,
			expectedStatusCodeVerify: 401,
			expectedStatusCode:       401,
			expectedResponseBody:     "mock",
		},
	}

	r := gin.New()
	auth := r.Group("/auth")
	{
		auth.POST("/sign-in", handlers.SignIn)
		auth.POST("/sign-up", handlers.SignUp)
	}
	api := r.Group("/api", handlers.UserIdentity)
	{
		api.POST("/info", handlers.GetInfo)
	}

	req, err := http.NewRequest("POST", "/auth/sign-up", bytes.NewBufferString(`{"username": "admin", "password": "admin"}`))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			req, err = http.NewRequest("POST", "/auth/sign-in", bytes.NewBufferString(testCase.signInBody))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			w = httptest.NewRecorder()
			r.ServeHTTP(w, req)
			require.Equal(t, testCase.expectedStatusCodeVerify, w.Code)
			token := strings.Split(w.Body.String(), ":")[1]
			token = token[1:]
			token = token[:len(token)-2]

			req, err = http.NewRequest("POST", "/api/info", bytes.NewBufferString(testCase.requestBody))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			req.Header.Set("Authorization", "Bearer "+token)
			w = httptest.NewRecorder()

			r.ServeHTTP(w, req)
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			if testCase.expectedResponseBody != "mock" {
				require.Equal(t, testCase.expectedResponseBody, w.Body.String())
			}
		})
	}

	_ = db.Close()
}

func TestHandler_SendCoin(t *testing.T) {
	if err := initConfig(); err != nil {
		slog.Error("Error: init configs")
		os.Exit(1)
	}

	db, err := postgres.NewDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: "qwerty",
	})
	if err != nil {
		slog.Error("Error: failed to init db connection ")
		os.Exit(1)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	testTable := []struct {
		name                     string
		signInBody               string
		requestBody              string
		expectedStatusCodeVerify int
		expectedStatusCode       int
		expectedResponseBody     string
	}{
		{
			name:                     "SendCoins Ok",
			signInBody:               `{"username": "admin", "password": "admin"}`,
			requestBody:              `{"received_id": 2, "amount": 1}`,
			expectedStatusCodeVerify: 200,
			expectedStatusCode:       200,
			expectedResponseBody:     ``,
		},
		{
			name:                     "Bad password token",
			signInBody:               `{"username": "admin", "password": "admin1"}`,
			requestBody:              ``,
			expectedStatusCodeVerify: 401,
			expectedStatusCode:       401,
			expectedResponseBody:     "mock",
		},
	}

	r := gin.New()
	auth := r.Group("/auth")
	{
		auth.POST("/sign-in", handlers.SignIn)
		auth.POST("/sign-up", handlers.SignUp)
	}
	api := r.Group("/api", handlers.UserIdentity)
	{
		api.POST("/sendCoin", handlers.SendCoin)
	}

	req, err := http.NewRequest("POST", "/auth/sign-up", bytes.NewBufferString(`{"username": "admin", "password": "admin"}`))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	req, err = http.NewRequest("POST", "/auth/sign-up", bytes.NewBufferString(`{"username": "admin1", "password": "admin1"}`))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			req, err = http.NewRequest("POST", "/auth/sign-in", bytes.NewBufferString(testCase.signInBody))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			w = httptest.NewRecorder()
			r.ServeHTTP(w, req)
			require.Equal(t, testCase.expectedStatusCodeVerify, w.Code)
			token := strings.Split(w.Body.String(), ":")[1]
			token = token[1:]
			token = token[:len(token)-2]

			req, err = http.NewRequest("POST", "/api/sendCoin", bytes.NewBufferString(testCase.requestBody))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			req.Header.Set("Authorization", "Bearer "+token)
			w = httptest.NewRecorder()

			r.ServeHTTP(w, req)
			require.Equal(t, testCase.expectedStatusCode, w.Code)
			if testCase.expectedResponseBody != "mock" {
				require.Equal(t, testCase.expectedResponseBody, w.Body.String())
			}
		})
	}

	_ = db.Close()
}
