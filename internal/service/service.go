package service

import "github.com/quanergyO/avito_assingment/internal/repository"

type Service struct {
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
