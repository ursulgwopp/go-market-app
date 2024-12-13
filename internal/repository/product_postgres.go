package repository

import (
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

func (r *ProductPostgres) AddProduct(product models.ProductRequest) error {
	query := `INSERT INTO products (name, description, price, quantity) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, product.Name, product.Description, product.Price, product.Quantity)

	return err
}

func (r *ProductPostgres) GetAllProducts() ([]models.Product, error) {
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
	query := `SELECT name, description, price, quantity FROM products WHERE id = $1`
	err := r.db.Get(&product, query, id)

	return product, err
}

func (r *ProductPostgres) UpdateProduct(id int, product models.ProductRequest) error {
	query := `UPDATE products SET name = $1, description = $2, price = $3, quantity = $4 WHERE id = $5`
	_, err := r.db.Exec(query, product.Name, product.Description, product.Price, product.Quantity, id)

	return err
}

func (r *ProductPostgres) DeleteProduct(id int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.Exec(query, id)

	return err
}
