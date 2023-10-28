package services

import (
	"context"

	"github.com/kumin/AndPadDating/entities"
	"github.com/kumin/AndPadDating/repos"
)

type UserService struct {
	userRepo repos.UserRepo
}

func NewUserService(
	userRepo repos.UserRepo,
) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) CreateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	return u.userRepo.CreateOne(ctx, user)
}

func (u *UserService) GetUser(ctx context.Context, id int64) (*entities.User, error) {
	return u.userRepo.GetOne(ctx, id)
}

func (u *UserService) UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	return u.userRepo.UpdateOne(ctx, user)
}

func (u *UserService) DeleteUser(ctx context.Context, id int64) error {
	return u.userRepo.DeleteOne(ctx, id)
}
