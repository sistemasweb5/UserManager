package service

import (
	"service/rest-api/internal/core/domain"
	"service/rest-api/internal/port/in"
	"service/rest-api/internal/port/out"

	"github.com/google/uuid"
)

type ClientService struct {
	repo out.ClientRepository
}

func NewClientService(repo out.ClientRepository) in.ClientService {
	return &ClientService{repo: repo}
}

func (p *ClientService) GetAll() (*[]domain.Client, error) {
	return p.repo.GetAll()
}

func (p *ClientService) GetById(id *uuid.UUID) (*domain.Client, error) {
	return p.repo.GetById(id)
}

