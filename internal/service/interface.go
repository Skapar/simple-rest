package service

import (
	"context"

	"github.com/Skapar/simple-rest/internal/models/entities"
	pb "github.com/Skapar/simple-rest/proto"
)

type AuthService interface {
	RegisterUser(user *entities.User) (string, string, error)
}

type UserService interface {
	GetUserById(userID int64) (*entities.User, error)
	UpdateUser(userID int64, user *entities.User) error
	DeleteUser(userID int64) (*entities.User, error)
	SoftDeleteUser(userID int64) (*entities.User, error)
}

type ProfileService interface {
	GetProfile(ctx context.Context, in *pb.GetProfileRequest) (*pb.GetProfileResponse, error)
	CreateProfile(ctx context.Context, in *pb.CreateProfileRequest) (*pb.CreateProfileResponse, error)
	UpdateProfile(ctx context.Context, in *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error)
	DeleteProfile(ctx context.Context, in *pb.DeleteProfileRequest) (*pb.DeleteProfileResponse, error)
}
