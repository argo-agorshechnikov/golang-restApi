package repository

import (
	"database/sql"

	"github.com/argo-agorshechnikov/golang-restApi/internal/models"

	_ "github.com/lib/pq"
)

type UserRep struct {
	db *sql.DB
}

func NewUserRep(connStr string) (*UserRep, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &UserRep{db: db}, nil
}

func (r *UserRep) CreateUserRep(user *models.User) error {

	query := `INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)`

	_, err := r.db.Exec(query, user.ID, user.Name, user.Email, user.Password)
	return err
}

func (r *UserRep) GetUserByID(id string) (*models.User, error) {
	var user models.User
	query := `SELECT id, name, email FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
