package service

import (
	"github.com/Skapar/simple-rest/internal/models/entities"
)

type AuthService interface {
	RegisterUser(user *entities.User) (string, string, error)
}

type UserService interface {
	UpdateUser(userID int64, user *entities.User) error
	DeleteUser(userID int64) error
	SoftDeleteUser(userID int64) error
}
