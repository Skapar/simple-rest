package repository

import (
	"github.com/Skapar/simple-rest/internal/models/entities"
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
	return r.db.Create(user).Error
}
