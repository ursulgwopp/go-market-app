package service

import (
	"errors"

	"github.com/ursulgwopp/go-market-app/internal/models"
	"github.com/ursulgwopp/go-market-app/internal/repository"
)

type PurchaseService struct {
	repo repository.Purchase
}

func NewPurchaseService(repo repository.Purchase) *PurchaseService {
	return &PurchaseService{repo: repo}
}

func (s *PurchaseService) MakePurchase(userId int, productId int, quantity int) (int, error) {
	if quantity < 1 {
		return -1, errors.New("invalid quantity")
	}

	return s.repo.MakePurchase(userId, productId, quantity)
}

func (s *PurchaseService) GetUserPurchases(userId int) ([]models.Purchase, error) {
	return s.repo.GetUserPurchases(userId)
}

func (s *PurchaseService) GetProductPurchases(productId int) ([]models.Purchase, error) {
	return s.repo.GetProductPurchases(productId)
}
