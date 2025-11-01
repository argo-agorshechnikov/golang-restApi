package service

import (
	"errors"

	"github.com/argo-agorshechnikov/golang-restApi/internal/models"
	"github.com/argo-agorshechnikov/golang-restApi/internal/repository"
)

type ProductService struct {
	productRep *repository.ProductRep
}

func NewProductService (productRep *repository.ProductRep) *ProductService {
	return &ProductService{productRep: productRep}
}

func (p *ProductService) CreateProductService (product *models.Product) error{

	if product.ID == "" || product.Name == "" || product.Price == "" || product.Description == "" {
		return errors.New("id, name, price or desc cannot be empty")
	}

	return p.productRep.CreateProductRep(product)
}

func (p *ProductService) GetProductByIdService(id string) (*models.Product, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}


	return p.productRep.GetProductById(id)
}