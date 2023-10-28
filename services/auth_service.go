package services

import "context"

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) Login(ctx context.Context, phone string) (string, error) {
	return "", nil
}

func (a *AuthService) ValidateToken(ctx context.Context, token string) bool {
	return false
}
