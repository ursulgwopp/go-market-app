package service

import (
	"github.com/ursulgwopp/market-api/internal/models"
	"github.com/ursulgwopp/market-api/internal/repository"
)

type Authorization interface {
	SignUp(req models.SignUpRequest) (int, error)
	GenerateToken(req models.SignInRequest) (string, error)
	ParseToken(token string) (int, error)
}

type User interface {
	GetUserByID(userId int) (models.User, error)
	ListUsers() ([]models.User, error)
	DeleteUser(userId int) error
}

type Profile interface {
	GetProfile(userId int) (models.User, error)
	Deposit(userId int, amount int) error
}

type Product interface {
	AddProduct(ownerId int, req models.ProductRequest) error
	ListProducts() ([]models.Product, error)
	GetProductByID(productId int) (models.Product, error)
	UpdateProduct(ownerId int, productId int, req models.ProductRequest) error
	DeleteProduct(ownerId int, productId int) error
}

type Purchase interface {
	MakePurchase(userId int, productId int, quantity int) (int, error)
	GetUserPurchases(userId int) ([]models.Purchase, error)
	GetProductPurchases(productId int) ([]models.Purchase, error)
}

type Service struct {
	Authorization
	User
	Profile
	Product
	Purchase
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User:          NewUserService(repos.User),
		Profile:       NewProfileService(repos.Profile),
		Product:       NewProductService(repos.Product),
		Purchase:      NewPurchaseService(repos.Purchase),
	}
}
