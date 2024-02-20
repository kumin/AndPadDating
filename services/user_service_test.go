//go:build unit
// +build unit

package services

import (
	"context"
	"testing"

	"github.com/kumin/BityDating/entities"
	mocks_data "github.com/kumin/BityDating/mocks/data"
	mocks_repo "github.com/kumin/BityDating/mocks/repos"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserServiceSuiteTest struct {
	suite.Suite
	userService *UserService
	userRepo    *mocks_repo.UserRepo
	fileRepo    *mocks_repo.FileRepo
}

func (u *UserServiceSuiteTest) SetupTest() {
	u.userRepo = mocks_repo.NewUserRepo(u.T())
	u.fileRepo = mocks_repo.NewFileRepo(u.T())
	u.userService = NewUserService(u.userRepo, u.fileRepo)
}

func (u *UserServiceSuiteTest) TestUserService_Create() {
	u.userRepo.On("CreateOne", mock.AnythingOfType("context.backgroundCtx"), mock.AnythingOfType("*entities.User")).
		Return(func(ctx context.Context, user *entities.User) (*entities.User, error) {
			return mocks_data.Users[0], nil
		})
	user, err := u.userService.CreateUser(context.Background(), mocks_data.Users[0])
	u.Equal(mocks_data.Users[0].Id, user.Id)
	u.Nil(err)
}

func (u *UserServiceSuiteTest) TestUserService_Get() {
	u.userRepo.On("GetOne", mock.AnythingOfType("context.backgroundCtx"), mock.AnythingOfType("int64")).
		Return(func(ctx context.Context, id int64) (*entities.User, error) {
			return mocks_data.Users[0], nil
		})
	user, err := u.userService.GetUser(context.Background(), 1)
	u.Nil(err)
	u.Equal(int64(1), user.Id)
}

func (u *UserServiceSuiteTest) TestUserService_Update() {
	u.userRepo.On("UpdateOne", mock.AnythingOfType("context.backgroundCtx"), mock.AnythingOfType("*entities.User")).
		Return(func(ctx context.Context, user *entities.User) (*entities.User, error) {
			return mocks_data.Users[0], nil
		})
	_, err := u.userService.UpdateUser(context.Background(), mocks_data.Users[0])
	u.Nil(err)
}

func (u *UserServiceSuiteTest) TestUserService_Delete() {
	u.userRepo.On("DeleteOne", mock.AnythingOfType("context.backgroundCtx"), mock.AnythingOfType("int64")).
		Return(func(ctx context.Context, id int64) error {
			return nil
		})
	err := u.userService.DeleteUser(context.Background(), int64(1))
	u.Nil(err)
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceSuiteTest))
}
