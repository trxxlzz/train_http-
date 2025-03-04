package service

import (
	"errors"
	"time"
	"training/internal/models"
	"training/internal/repository"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id int) (*models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *models.User) error {
	if user.Name == "" || user.Email == "" || user.Age <= 0 {
		return errors.New("name, email and age are required and age must be positive")
	}

	user.CreatedAt = time.Now()
	return s.repo.CreateUser(user)
}

func (s *userService) GetUserByID(id int) (*models.User, error) {
	return s.repo.GetUserByID(id)
}
