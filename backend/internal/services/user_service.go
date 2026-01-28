package services

import (
	"errors"

	"github.com/Geovannisouza23/go-user-api/internal/models"
	"github.com/Geovannisouza23/go-user-api/internal/repositories"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.UserRepo.FindAll()
}

func (s *UserService) GetByID(id uint) (*models.User, error) {
	return s.UserRepo.FindByID(id)
}

func (s *UserService) Update(id uint, name string) (*models.User, error) {
	user, err := s.UserRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	user.Name = name
	err = s.UserRepo.Update(user)
	return user, err
}

func (s *UserService) Delete(id uint) error {
	return s.UserRepo.Delete(id)
}
