package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/argo-agorshechnikov/golang-restApi/internal/models"
	"github.com/argo-agorshechnikov/golang-restApi/internal/service"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler (productService *service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}


func (p *ProductHandler) CreateProductHandler (w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Incorrect data format", http.StatusBadRequest)
		return
	}

	err = p.productService.CreateProductService(&product)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}
