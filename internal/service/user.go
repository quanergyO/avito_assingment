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
	return s.repo.User.GetUserInfo(userID)
}

func (s *UserService) SendCoins(senderID, receiverID int, amount int) error {
	return s.repo.User.SendCoins(senderID, receiverID, amount)
}

func (s *UserService) BuyItem(userID int, itemName string) error {
	if !s.isItemExists(itemName) {
		return fmt.Errorf("item is not exists")
	}

	return s.repo.User.BuyItem(userID, itemName)
}

func (s *UserService) isItemExists(itemName string) bool {
	items := [...]string{"t-shirt", "cup", "book", "pen", "powerbank", "hoody", "umbrella", "socks", "wallet", "pink-hooby"} // TODO should be static member
	for _, item := range items {
		if item == itemName {
			return true
		}
	}
	return false
}
