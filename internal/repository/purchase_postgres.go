package repository

import (
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

func (r *PurchasePostgres) MakePurchase(purchase models.Purchase, quantity int) (int, error) {
	var productQuantity int
	query := `SELECT quantity FROM products WHERE id = $1`
	err := r.db.QueryRow(query, purchase.ProductId).Scan(&productQuantity)
	if err != nil {
		return -1, err
	}

	if productQuantity-quantity < 0 {
		return -1, fmt.Errorf("not enough products")
	}

	query = `UPDATE products SET quantity = $1 WHERE id = $2`
	_, err = r.db.Exec(query, productQuantity-quantity, purchase.ProductId)
	if err != nil {
		return -1, err
	}

	var id int
	query = `INSERT INTO purchases (user_id, product_id, quantity, timestamp) VALUES ($1, $2, $3, $4) RETURNING id`
	row := r.db.QueryRow(query, purchase.UserId, purchase.ProductId, quantity, (time.Now().String())[:19])
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
		if err := rows.Scan(&purchase.Id, &purchase.UserId, &purchase.ProductId, &purchase.Quantity, &purchase.Timestamp); err != nil {
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
		if err := rows.Scan(&purchase.Id, &purchase.UserId, &purchase.ProductId, &purchase.Quantity, &purchase.Timestamp); err != nil {
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
