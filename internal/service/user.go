package service

import (
	"github.com/ursulgwopp/market-api/internal/models"
	"github.com/ursulgwopp/market-api/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(userId int) (models.User, error) {
	return s.repo.GetUserByID(userId)
}

func (s *UserService) ListUsers() ([]models.User, error) {
	return s.repo.ListUsers()
}

func (s *UserService) DeleteUser(userId int) error {
	return s.repo.DeleteUser(userId)
}
