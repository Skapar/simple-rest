package repository

import (
	"github.com/Skapar/simple-rest/internal/models/entities"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *entities.User) error
}

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{db: db}
}

func (r *AuthRepositoryImpl) CreateUser(user *entities.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}
