package v2

import (
	"errors"
	"time"
	"training/internal/models"
	"training/internal/repository"
)

type UserServiceV2 interface {
	CreateUserV2(user *models.UserV2) error
	GetUserByIDV2(id int) (*models.UserV2, error)
}

type userServiceV2 struct {
	repo repository.UserRepository
}

func NewUserServiceV2(repo repository.UserRepository) UserServiceV2 {
	return &userServiceV2{repo: repo}
}

func (s *userServiceV2) CreateUserV2(user *models.UserV2) error {
	if user.FullName == "" || user.Email == "" || user.Age <= 0 {
		return errors.New("full name, email and age are required and age must be positive")
	}

	user.CreatedAt = time.Now()

	// Тут можно репозиторий адаптировать, если схема в БД отличается для v2
	return s.repo.CreateUserV2(user)
}

func (s *userServiceV2) GetUserByIDV2(id int) (*models.UserV2, error) {
	return s.repo.GetUserByIDV2(id)
}
