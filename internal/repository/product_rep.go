package repository

import (
	"database/sql"

	"github.com/argo-agorshechnikov/golang-restApi/internal/models"
)

type ProductRep struct {
	db *sql.DB
}

func NewProductRep(connStr string) (*ProductRep, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &ProductRep{db: db}, nil
}

func (r *ProductRep) CreateProductRep(product *models.Product) error {

	query := "INSERT INTO products(id, name, price, description) VALUES ($1, $2, $3, $4)"

	_, err := r.db.Exec(query, product.ID, product.Name, product.Description)
	return err
}

func (r *ProductRep) GetProductById(id string) (*models.Product, error) {
	var product models.Product

	query := "SELECT id, name, desctiptions FROM products WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description)
	if err != nil {
		return nil, err
	}

	return &product, nil

}
