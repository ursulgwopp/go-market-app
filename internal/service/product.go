package service

import (
	"github.com/ursulgwopp/go-market-app/internal/models"
	"github.com/ursulgwopp/go-market-app/internal/repository"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) AddProduct(ownerId int, req models.ProductRequest) error {
	if err := validateProduct(req); err != nil {
		return err
	}

	return s.repo.AddProduct(ownerId, req)
}

func (s *ProductService) ListProducts() ([]models.Product, error) {
	return s.repo.ListProducts()
}

func (s *ProductService) GetProductByID(productId int) (models.Product, error) {
	return s.repo.GetProductByID(productId)
}

func (s *ProductService) UpdateProduct(ownerId int, productId int, req models.ProductRequest) error {
	if err := validateProduct(req); err != nil {
		return err
	}

	return s.repo.UpdateProduct(ownerId, productId, req)
}

func (s *ProductService) DeleteProduct(ownerId int, productId int) error {
	return s.repo.DeleteProduct(ownerId, productId)
}
