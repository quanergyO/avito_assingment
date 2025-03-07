package service

import (
	"github.com/quanergyO/avito_assingment/internal/repository"
	"github.com/quanergyO/avito_assingment/types"
)

//go:generate mockgen -source=service.go -destination=../mocks/mock_services.go -package=mocks
type Authorization interface {
	CreateUser(user types.SignInInput) (int, error)
	CheckAuthData(username, password string) (types.UserDAO, error)
	GenerateToken(user types.UserDAO) (string, error)
	ParserToken(accessToken string) (*types.TokenClaims, error)
}

type User interface {
	GetUserInfo(userID int) (types.UserInfo, error)
	SendCoins(senderID, receiverID int, amount int) error
	BuyItem(userID int, itemName string) error
}

type Service struct {
	Authorization
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		User:          NewUserService(repo),
	}
}
