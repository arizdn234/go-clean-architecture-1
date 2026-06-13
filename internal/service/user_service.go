package service

import (
	"errors"

	"github.com/arizdn234/go-clean-architecture-1/internal/domain"
	"github.com/arizdn234/go-clean-architecture-1/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(
	repo repository.UserRepository,
) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(name string) error {
	if name == "" {
		return errors.New("name required")
	}

	user := domain.User{
		Name: name,
	}

	return s.repo.Save(user)
}