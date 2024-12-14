package repository

import (
	"sort"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/ursulgwopp/go-market-app/internal/models"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetUserByID(userId int) (models.User, error) {
	var user models.User
	query := `SELECT username, email, COALESCE(product_list, '{}') FROM users WHERE id = $1`
	err := r.db.QueryRow(query, userId).Scan(
		&user.Username,
		&user.Email,
		pq.Array(&user.ProductList))
	user.Id = userId

	return user, err
}

func (r *UserPostgres) ListUsers() ([]models.User, error) {
	query := `SELECT id, username, email, COALESCE(product_list, '{}') FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var user models.User

		if err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			pq.Array(&user.ProductList),
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	sortById := func(i, j int) bool {
		return users[i].Id < users[j].Id
	}
	sort.Slice(users, sortById)

	return users, nil
}

func (r *UserPostgres) DeleteUser(userId int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, userId)
	if err != nil {
		return err
	}

	query = `DELETE FROM products WHERE owner_id = $1`
	_, err = r.db.Exec(query, userId)

	return err
}

func (r *UserPostgres) CheckUserExists(userId int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)"
	err := r.db.QueryRow(query, userId).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
