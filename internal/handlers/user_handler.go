package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/argo-agorshechnikov/golang-restApi/internal/models"
	"github.com/argo-agorshechnikov/golang-restApi/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{userService: us}
}


func (h *UserHandler) CreateUserHand(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Uncorrect data format", http.StatusBadRequest)
		return
	}

	err = h.userService.CreateUserService(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)


}

