package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ursulgwopp/go-market-app/internal/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(req models.SignUpRequest) (int, error) {
	var id int
	query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id`

	row := r.db.QueryRow(query, req.Username, req.Password, req.Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(req models.SignInRequest) (models.User, error) {
	var newUser models.User
	query := `SELECT id FROM users WHERE username = $1 AND password = $2`
	err := r.db.Get(&newUser, query, req.Username, req.Password)

	return newUser, err
}
