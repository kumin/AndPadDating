package services

import (
	"context"
	"mime/multipart"

	"github.com/kumin/BityDating/entities"
	"github.com/kumin/BityDating/repos"
)

type UserService struct {
	userRepo repos.UserRepo
	fileRepo repos.FileRepo
}

func NewUserService(
	userRepo repos.UserRepo,
	fileRepo repos.FileRepo,
) *UserService {
	return &UserService{
		userRepo: userRepo,
		fileRepo: fileRepo,
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

func (u *UserService) SetAvatar(ctx context.Context, file *multipart.FileHeader) (string, error) {
	fileEntity, err := ConvertMultipartToFile(file)
	if err != nil {
		return "", err
	}
	return u.fileRepo.UploadFile(ctx, fileEntity)
}
