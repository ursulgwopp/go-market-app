package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ursulgwopp/market-api/internal/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) SignUp(req models.SignUpRequest) (int, error) {
	var id int
	query := `INSERT INTO users (username, hash_password, email) VALUES ($1, $2, $3) RETURNING id`

	row := r.db.QueryRow(query, req.Username, req.Password, req.Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) SignIn(req models.SignInRequest) (int, error) {
	var user_id int
	query := `SELECT id FROM users WHERE username = $1 AND hash_password = $2`
	err := r.db.Get(&user_id, query, req.Username, req.Password)

	return user_id, err
}

func (r *AuthPostgres) CheckUsernameExists(username string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)"
	err := r.db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
