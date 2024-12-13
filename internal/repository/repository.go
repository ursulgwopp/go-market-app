package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ursulgwopp/go-market-app/internal/models"
)

type Authorization interface {
	CreateUser(req models.SignUpRequest) (int, error)
	GetUser(req models.SignInRequest) (models.User, error)
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

type Repository struct {
	Authorization
	User
	Product
	Purchase
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
		Product:       NewProductPostgres(db),
		Purchase:      NewPurchasePostgres(db),
	}
}
