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

func (s *AuthService) SignUp(user domain.UserSignUp) (string, error) {
	return s.authService.SignUp(user)
}

func (s *AuthService) ConfirmAccount(user domain.UserConfirmation) error {
	return s.authService.ConfirmAccount(user)
}

func (s *AuthService) GetUserIdByToken(accessToken string) (string, error) {
	return s.authService.GetUserIdByToken(accessToken)
}
