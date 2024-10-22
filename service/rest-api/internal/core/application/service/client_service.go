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

func (p *ClientService) GetAll() (*[]domain.ClientResponse, error) {
	var clientReponses []domain.ClientResponse
	clients, err := p.repo.GetAllClients()
	if err != nil {
		return nil, err
	}
	for _, client := range *clients {
		clientResponse, err := p.GetById(&client.Id)
		if err != nil {
			return nil, err
		}
		clientReponses = append(clientReponses, *clientResponse)
	}

	return &clientReponses, nil
}

func (p *ClientService) GetById(id *uuid.UUID) (*domain.ClientResponse, error) {
	client, err := p.repo.GetClientById(id)
	if err != nil {
		return nil, err
	}
	category, err := p.repo.GetCategoryById(&client.CategoryId)
	if err != nil {
		return nil, err
	}
	workSchedule, err := p.repo.GetWorkScheduleById(&client.WorkScheduleId)
	if err != nil {
		return nil, err
	}
	specialties, err := p.repo.GetSpecialitiesByClientId(&client.Id)
	if err != nil {
		return nil, err
	}

	return &domain.ClientResponse{
		Client:       *client,
		Category:     *category,
		WorkSchedule: *workSchedule,
		Specialties:  *specialties,
	}, nil
}
