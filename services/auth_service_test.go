//go:build unit
// +build unit

package services

import (
	"context"
	"testing"

	"github.com/kumin/AndPadDating/entities"
	mocks_data "github.com/kumin/AndPadDating/mocks/data"
	mocks_repo "github.com/kumin/AndPadDating/mocks/repos"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type AuthServiceTestSuit struct {
	suite.Suite
	authService *AuthService
	userRepo    *mocks_repo.UserRepo
}

func (a *AuthServiceTestSuit) SetupTest() {
	a.userRepo = mocks_repo.NewUserRepo(a.T())
	a.authService = NewAuthService(a.userRepo)
}

func (a *AuthServiceTestSuit) TestAuthService_RegisterUser_Happy() {
	a.userRepo.On("VerifyPhone", mock.AnythingOfType("context.backgroundCtx"), mock.AnythingOfType("string")).
		Return(func(ctx context.Context, phone string) (bool, error) {
			if phone == mocks_data.Users[0].Phone {
				return true, nil
			}
			return false, nil
		})
	a.userRepo.On("CreateOne", mock.AnythingOfType("context.backgroundCtx"), mock.AnythingOfType("*entities.User")).
		Return(func(ctx context.Context, user *entities.User) (*entities.User, error) {
			return mocks_data.Users[0], nil
		})
	registeredUser, err := a.authService.Register(context.Background(), mocks_data.Users[0])
	a.Nil(err)
	a.Equal(mocks_data.Users[0].Id, registeredUser.User.Id)
}

func (a *AuthServiceTestSuit) TestAuthService_RegisterUser_PhoneAlreadyExist() {
	a.userRepo.On("VerifyPhone", mock.AnythingOfType("context.backgroundCtx"), mock.AnythingOfType("string")).
		Return(func(ctx context.Context, phone string) (bool, error) {
			if phone == mocks_data.Users[0].Phone {
				return false, nil
			}
			return false, nil
		})
	registeredUser, err := a.authService.Register(context.Background(), mocks_data.Users[0])
	a.NotNil(err)
	a.Nil(registeredUser)
}

func (a *AuthServiceTestSuit) TestAuthService_Login_Happy() {
	a.userRepo.On("GetByPhone", mock.AnythingOfType("context.backgroundCtx"), mock.AnythingOfType("string")).
		Return(func(ctx context.Context, phone string) (*entities.User, error) {
			return mocks_data.Users[0], nil
		})
	registeredUser, err := a.authService.Login(context.Background(), mocks_data.Users[0].Phone)
	a.Nil(err)
	a.Equal(registeredUser.User.Id, mocks_data.Users[0].Id)

	ok := ValidateToken(registeredUser.Token)
	a.True(ok)
}

func (a *AuthServiceTestSuit) TestAuthService_Login_AccountNotExist() {
	a.userRepo.On("GetByPhone", mock.AnythingOfType("context.backgroundCtx"), mock.AnythingOfType("string")).
		Return(func(ctx context.Context, phone string) (*entities.User, error) {
			return nil, gorm.ErrRecordNotFound
		})
	registeredUser, err := a.authService.Login(context.Background(), mocks_data.Users[0].Phone)
	a.NotNil(err)
	a.Nil(registeredUser)
}

func TestAuthService(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuit))
}
