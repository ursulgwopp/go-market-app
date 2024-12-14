package service

import (
	"github.com/ursulgwopp/go-market-app/internal/models"
	"github.com/ursulgwopp/go-market-app/internal/repository"
)

type ProfileService struct {
	repo repository.Profile
}

func NewProfileService(repo repository.Profile) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) GetProfile(userId int) (models.User, error) {
	return s.repo.GetProfile(userId)
}

func (s *ProfileService) Deposit(userId int, amount int) error {
	return s.repo.Deposit(userId, amount)
}
