package repository

import (
	"errors"
	"sort"

	"github.com/jmoiron/sqlx"
	"github.com/ursulgwopp/go-market-app/internal/models"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) AddProduct(ownerId int, req models.ProductRequest) error {
	query := `INSERT INTO products (name, description, price, quantity, owner_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	row := r.db.QueryRow(query, req.Name, req.Description, req.Price, req.Quantity, ownerId)

	var productId int
	if err := row.Scan(&productId); err != nil {
		return err
	}

	query = `UPDATE users SET product_list = array_append(product_list, $1) WHERE id = $2`
	_, err := r.db.Exec(query, productId, ownerId)

	return err
}

func (r *ProductPostgres) ListProducts() ([]models.Product, error) {
	query := `SELECT * FROM products`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var products []models.Product
	for rows.Next() {
		var product models.Product

		if err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Quantity,
			&product.OwnerId,
		); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	sortById := func(i, j int) bool {
		return products[i].Id < products[j].Id
	}
	sort.Slice(products, sortById)

	return products, nil
}

func (r *ProductPostgres) GetProductByID(id int) (models.Product, error) {
	var product models.Product
	product.Id = id
	query := `SELECT name, description, price, quantity, owner_id FROM products WHERE id = $1`
	err := r.db.Get(&product, query, id)

	return product, err
}

func (r *ProductPostgres) UpdateProduct(ownerId, productId int, req models.ProductRequest) error {
	var dbOwnerId int
	query := `SELECT owner_id FROM products WHERE id = $1`
	err := r.db.QueryRow(query, productId).Scan(&dbOwnerId)
	if err != nil {
		return err
	}

	if dbOwnerId != ownerId {
		return errors.New("no rights")
	}

	query = `UPDATE products SET name = $1, description = $2, price = $3, quantity = $4 WHERE id = $5`
	_, err = r.db.Exec(query, req.Name, req.Description, req.Price, req.Quantity, productId)

	return err
}

func (r *ProductPostgres) DeleteProduct(ownerId, productId int) error {
	var dbOwnerId int
	query := `SELECT owner_id FROM products WHERE id = $1`
	err := r.db.QueryRow(query, productId).Scan(&dbOwnerId)
	if err != nil {
		return err
	}

	if dbOwnerId != ownerId {
		return errors.New("no rights")
	}

	query = `DELETE FROM products WHERE id = $1`
	_, err = r.db.Exec(query, productId)
	if err != nil {
		return err
	}

	query = `UPDATE users SET product_list = array_remove(product_list, $1) WHERE id = $2`
	_, err = r.db.Exec(query, productId, ownerId)

	return err
}
