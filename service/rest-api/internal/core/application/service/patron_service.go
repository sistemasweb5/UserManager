package service

import (
	"service/rest-api/internal/core/domain"
	"service/rest-api/internal/port/in"
	"service/rest-api/internal/port/out"
)

type UserService struct {
	repo out.UserRepository
}

func NewUserService(repo out.UserRepository) in.UserService {
	return &UserService{repo: repo}
}

func (p *UserService) GetAll() (*[]domain.User, error) {
	return p.repo.GetAll()
}

