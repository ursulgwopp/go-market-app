package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ursulgwopp/go-market-app/internal/models"
)

type Authorization interface {
	SignUp(req models.SignUpRequest) (int, error)
	SignIn(req models.SignInRequest) (int, error)
	CheckUsernameExists(username string) (bool, error)
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

type Repository struct {
	Authorization
	User
	Profile
	Product
	Purchase
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
		Profile:       NewProfilePostgres(db),
		Product:       NewProductPostgres(db),
		Purchase:      NewPurchasePostgres(db),
	}
}
