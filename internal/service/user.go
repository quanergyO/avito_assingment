package service

import (
	"fmt"

	"github.com/quanergyO/avito_assingment/internal/repository"
	"github.com/quanergyO/avito_assingment/types"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUserInfo(userID int) (types.UserInfo, error) {
	return types.UserInfo{}, fmt.Errorf("Not implemented")
}

func (s *UserService) SendCoins(senderID, receiverID int, amount int) error {
	return fmt.Errorf("Not implemented")
}

func (s *UserService) BuyItem(userID int, itemName string) error {
	return fmt.Errorf("Not implemented")
}
