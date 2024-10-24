package service

import (
	"service/rest-api/internal/core/domain"
	"service/rest-api/internal/port/in"
)

type AuthService struct {
	authService in.AuthService
}

func NewAuthService(authService in.AuthService) *AuthService {
	return &AuthService{
		authService: authService,
	}
}

func (s *AuthService) SignIn(user domain.UserLogin) (string, error) {
	return s.authService.SignIn(user)
}

func (s *AuthService) Logout(accessToken string) error {
	return s.authService.Logout(accessToken)
}
