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

func (s *ProductService) AddProduct(product models.ProductRequest) error {
	return s.repo.AddProduct(product)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAllProducts()
}

func (s *ProductService) GetProductByID(id int) (models.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s *ProductService) UpdateProduct(id int, product models.ProductRequest) error {
	return s.repo.UpdateProduct(id, product)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.repo.DeleteProduct(id)
}
