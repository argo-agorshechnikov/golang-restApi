package service

import (
	"errors"

	"github.com/argo-agorshechnikov/golang-restApi/internal/models"
	"github.com/argo-agorshechnikov/golang-restApi/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRep
}

func NewUserService(repo repository.UserRep) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUserService(user *models.User) error {
	if user.Email == "" || user.Password == "" || user.Name == "" {
		return errors.New("name, email or password cannot be empty")
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashPass)

	return s.repo.CreateUserRep(user)
}

func (s *UserService) GetUser(id string) (*models.User, error) {
	return s.repo.GetUserByID(id)
}
