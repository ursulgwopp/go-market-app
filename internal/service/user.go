package service

import (
	"github.com/ursulgwopp/go-market-app/internal/models"
	"github.com/ursulgwopp/go-market-app/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(id int) (models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}
