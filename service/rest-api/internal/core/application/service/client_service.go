package service

import (
	"context"
	"log"
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

func (p *ClientService) GetAll(ctx context.Context) (*[]domain.ClientResponse, error) {
	var clientReponses []domain.ClientResponse
	clients, err := p.repo.GetAllClients(ctx)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	for _, client := range *clients {
		clientResponse, err := p.GetById(ctx, &client.Id)
		if err != nil {
			log.Printf("Error: %v", err)
			return nil, err
		}
		clientReponses = append(clientReponses, *clientResponse)
	}

	return &clientReponses, nil
}

func (p *ClientService) GetById(ctx context.Context, id *uuid.UUID) (*domain.ClientResponse, error) {
	client, err := p.repo.GetClientById(ctx, id)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	category, err := p.repo.GetCategoryById(ctx, &client.CategoryId)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	workSchedule, err := p.repo.GetWorkScheduleById(ctx, &client.WorkScheduleId)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	specialties, err := p.repo.GetSpecialitiesByClientId(ctx, &client.Id)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}

	return &domain.ClientResponse{
		Client:       *client,
		Category:     *category,
		WorkSchedule: *workSchedule,
		Specialties:  *specialties,
	}, nil
}
