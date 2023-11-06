package services

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kumin/AndPadDating/configs"
	"github.com/kumin/AndPadDating/entities"
	"github.com/kumin/AndPadDating/erroz"
	"github.com/kumin/AndPadDating/repos"
	"github.com/rs/zerolog/log"
)

type AuthService struct {
	userRepo repos.UserRepo
}

func NewAuthService(
	userRepo repos.UserRepo,
) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (a *AuthService) Register(ctx context.Context, user *entities.User) (*entities.RegisteredUser, error) {
	if ok, err := a.userRepo.VerifyPhone(ctx, user.Phone); !ok || err != nil {
		if err != nil {
			return nil, err
		}
		return nil, erroz.ErrPhoneAlreadyUsed
	}
	registeredUser, err := a.userRepo.CreateOne(ctx, user)
	if err != nil {
		return nil, err
	}
	token, err := a.genToken(registeredUser)
	if err != nil {
		return nil, err
	}

	return &entities.RegisteredUser{
		Token: token,
		User:  registeredUser,
	}, nil
}

func (a *AuthService) Login(ctx context.Context, phone string) (*entities.RegisteredUser, error) {
	user, err := a.userRepo.GetByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	token, err := a.genToken(user)
	if err != nil {
		return nil, err
	}

	return &entities.RegisteredUser{
		Token: token,
		User:  user,
	}, nil
}

func ValidateToken(token string) bool {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("[auth services] unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(configs.SecretKey), nil
	})
	if err != nil {
		log.Error().Msg(fmt.Sprintf("[auth services] %s", err.Error()))
		return false
	}

	return jwtToken.Valid
}

func (a *AuthService) genToken(user *entities.User) (string, error) {
	claims := &entities.Claims{
		Phone:  user.Phone,
		UserId: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(configs.TokenExpiredTime)),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(configs.SecretKey))
}
