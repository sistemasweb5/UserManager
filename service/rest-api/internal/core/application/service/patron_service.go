package service

import (
	"service/rest-api/internal/core/domain"
	"service/rest-api/internal/port/in"
	"service/rest-api/internal/port/out"
	"rest-api/internal/adapter/http"
)

type UserService struct {
	repo out.UserRepository
}

type AuthService struct {
    cognito http.CognitoInterface
}

func NewUserService(repo out.UserRepository) in.UserService {
	return &UserService{repo: repo}
}

func (p *UserService) GetAll() (*[]domain.User, error) {
	return p.repo.GetAll()
}

func NewAuthService(cognito http.CognitoInterface) *AuthService {
    return &AuthService{cognito: cognito}
}

func (s *AuthService) Login(user *http.UserLogin) (string, error) {
    return s.cognito.SignIn(user)
}
