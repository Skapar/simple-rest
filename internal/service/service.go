package service

import (
	"github.com/Skapar/simple-rest/internal/models/entities"
	"github.com/Skapar/simple-rest/internal/repository"
)

type AuthService interface {
	RegisterUser(user *entities.User) error
}

type AuthServiceImpl struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &AuthServiceImpl{repo: repo}
}

func (s *AuthServiceImpl) RegisterUser(user *entities.User) error {
	return s.repo.CreateUser(user)
}
