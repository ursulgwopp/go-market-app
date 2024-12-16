package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/ursulgwopp/market-api/internal/models"
)

type ProfilePostgres struct {
	db *sqlx.DB
}

func NewProfilePostgres(db *sqlx.DB) *ProfilePostgres {
	return &ProfilePostgres{db: db}
}

func (r *ProfilePostgres) GetProfile(userId int) (models.User, error) {
	var user models.User
	query := `SELECT username, email, balance, COALESCE(product_list, '{}') FROM users WHERE id = $1`
	err := r.db.QueryRow(query, userId).Scan(
		&user.Username,
		&user.Email,
		&user.Balance,
		pq.Array(&user.ProductList))
	user.Id = userId

	return user, err
}

func (r *ProfilePostgres) Deposit(userId int, amount int) error {
	query := `UPDATE users SET balance = balance + $1 WHERE id = $2`
	_, err := r.db.Exec(query, amount, userId)

	return err
}
