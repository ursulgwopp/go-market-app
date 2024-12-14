package repository

import (
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ursulgwopp/go-market-app/internal/models"
)

type PurchasePostgres struct {
	db *sqlx.DB
}

func NewPurchasePostgres(db *sqlx.DB) *PurchasePostgres {
	return &PurchasePostgres{db: db}
}

func (r *PurchasePostgres) MakePurchase(userId int, productId int, quantity int) (int, error) {
	var productPrice, productQuantity int
	query := `SELECT price, quantity FROM products WHERE id = $1`
	err := r.db.QueryRow(query, productId).Scan(&productPrice, &productQuantity)
	if err != nil {
		return -1, err
	}

	if productQuantity-quantity < 0 {
		return -1, fmt.Errorf("not enough products")
	}

	var userBalance int
	query = `SELECT balance FROM users WHERE id = $1`
	err = r.db.QueryRow(query, userId).Scan(&userBalance)
	if err != nil {
		return -1, err
	}

	if userBalance-(productId*quantity) < 0 {
		return -1, errors.New("u fucking broke")
	}

	query = `UPDATE products SET quantity = $1 WHERE id = $2`
	_, err = r.db.Exec(query, productQuantity-quantity, productId)
	if err != nil {
		return -1, err
	}

	query = `UPDATE users SET balance = $1 WHERE id = $2`
	_, err = r.db.Exec(query, userBalance-(quantity*productPrice), userId)
	if err != nil {
		return -1, err
	}

	var id int
	query = `INSERT INTO purchases (user_id, product_id, quantity, timestamp, cost) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	row := r.db.QueryRow(query, userId, productId, quantity, (time.Now().String())[:19], quantity*productPrice)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *PurchasePostgres) GetUserPurchases(id int) ([]models.Purchase, error) {
	query := `SELECT * FROM purchases WHERE user_id = $1`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	purchases := []models.Purchase{}
	for rows.Next() {
		purchase := models.Purchase{}
		if err := rows.Scan(&purchase.Id, &purchase.UserId, &purchase.ProductId, &purchase.Quantity, &purchase.Timestamp, &purchase.Cost); err != nil {
			return nil, err
		}

		purchases = append(purchases, purchase)
	}

	sortById := func(i, j int) bool {
		return purchases[i].Id < purchases[j].Id
	}
	sort.Slice(purchases, sortById)

	return purchases, nil
}

func (r *PurchasePostgres) GetProductPurchases(id int) ([]models.Purchase, error) {
	query := `SELECT * FROM purchases WHERE product_id = $1`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	purchases := []models.Purchase{}
	for rows.Next() {
		purchase := models.Purchase{}
		if err := rows.Scan(&purchase.Id, &purchase.UserId, &purchase.ProductId, &purchase.Quantity, &purchase.Timestamp, &purchase.Cost); err != nil {
			return nil, err
		}

		purchases = append(purchases, purchase)
	}

	sortById := func(i, j int) bool {
		return purchases[i].Id < purchases[j].Id
	}
	sort.Slice(purchases, sortById)

	return purchases, nil
}
