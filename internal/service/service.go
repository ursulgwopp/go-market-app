package service

import (
	"github.com/ursulgwopp/go-market-app/internal/models"
	"github.com/ursulgwopp/go-market-app/internal/repository"
)

type Authorization interface {
	CreateUser(req models.SignUpRequest) (int, error)
	GenerateToken(req models.SignInRequest) (string, error)
	ParseToken(token string) (int, error)
}

type User interface {
	GetUserByID(id int) (models.User, error)
	GetAllUsers() ([]models.User, error)
}

type Product interface {
	AddProduct(product models.ProductRequest) error
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id int) (models.Product, error)
	UpdateProduct(id int, product models.ProductRequest) error
	DeleteProduct(id int) error
}

type Purchase interface {
	MakePurchase(purchase models.Purchase, quantity int) (int, error)
	GetUserPurchases(id int) ([]models.Purchase, error)
	GetProductPurchases(id int) ([]models.Purchase, error)
}

type Service struct {
	Authorization
	User
	Product
	Purchase
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User:          NewUserService(repos.User),
		Product:       NewProductService(repos.Product),
		Purchase:      NewPurchaseService(repos.Purchase),
	}
}
