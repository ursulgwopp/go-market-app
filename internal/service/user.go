package service

import (
	"errors"

	"github.com/ursulgwopp/go-market-app/internal/models"
	"github.com/ursulgwopp/go-market-app/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(userId int) (models.User, error) {
	exists, err := s.repo.CheckUserExists(userId)
	if err != nil {
		return models.User{}, err
	}

	if !exists {
		return models.User{}, errors.New("invalid id")
	}

	return s.repo.GetUserByID(userId)
}

func (s *UserService) ListUsers() ([]models.User, error) {
	return s.repo.ListUsers()
}

func (s *UserService) DeleteUser(userId int) error {
	exists, err := s.repo.CheckUserExists(userId)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("invalid id")
	}

	return s.repo.DeleteUser(userId)
}
