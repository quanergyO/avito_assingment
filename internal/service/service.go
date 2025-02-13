package service

import (
	"github.com/quanergyO/avito_assingment/internal/repository"
	"github.com/quanergyO/avito_assingment/types"
)

type Authorization interface {
	CreateUser(user types.UserType) (int, error)
	CheckAuthData(username, password string) (types.UserType, error)
	GenerateToken(user types.UserType) (string, error)
	ParserToken(accessToken string) (*types.TokenClaims, error)
}

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
	}
}
