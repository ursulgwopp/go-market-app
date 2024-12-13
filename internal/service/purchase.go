package service

import (
	"github.com/ursulgwopp/go-market-app/internal/models"
	"github.com/ursulgwopp/go-market-app/internal/repository"
)

type PurchaseService struct {
	repo repository.Purchase
}

func NewPurchaseService(repo repository.Purchase) *PurchaseService {
	return &PurchaseService{repo: repo}
}

func (s *PurchaseService) MakePurchase(purchase models.Purchase, quantity int) (int, error) {
	return s.repo.MakePurchase(purchase, quantity)
}

func (s *PurchaseService) GetUserPurchases(id int) ([]models.Purchase, error) {
	return s.repo.GetUserPurchases(id)
}

func (s *PurchaseService) GetProductPurchases(id int) ([]models.Purchase, error) {
	return s.repo.GetProductPurchases(id)
}
