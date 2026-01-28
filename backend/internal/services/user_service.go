package services

import (
	"errors"

	"github.com/Geovannisouza23/go-user-api/internal/models"
	"github.com/Geovannisouza23/go-user-api/internal/repositories"
	"gorm.io/gorm"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.UserRepo.GetAll()
}

func (s *UserService) GetByID(id uint) (*models.User, error) {
	return s.UserRepo.GetByID(id)
}

func (s *UserService) Update(id uint, name string) (*models.User, error) {
	user, err := s.UserRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	user.Name = name

	if err := s.UserRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Delete(id uint) error {
	return s.UserRepo.Delete(id)
}
