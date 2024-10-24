package service

import (
	"service/rest-api/internal/core/domain"
	"service/rest-api/internal/port/out"
)

type AuthService struct {
	repo out.AuthRepository
}

func NewAuthService(repo out.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp(user *domain.User) error {
	return s.repo.SignUp(user)
}

func (s *AuthService) ConfirmAccount(confirmation *domain.UserConfirmation) error {
	return s.repo.ConfirmAccount(confirmation)
}
