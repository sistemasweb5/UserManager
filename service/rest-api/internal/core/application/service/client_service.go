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

	workSchedule := new(domain.WorkSchedule)
	specialties := new([]domain.Specialty)
	workSchedule, err = p.repo.GetWorkScheduleById(ctx, &client.WorkScheduleId)
	if err != nil {
		log.Printf("Error: %v", err)
		workSchedule = &domain.WorkSchedule{} 
	}
	specialties, err = p.repo.GetSpecialitiesByClientId(ctx, &client.Id)
	if err != nil {
		log.Printf("Error: %v", err)
		specialties = &[]domain.Specialty{}
	}

	return &domain.ClientResponse{
		Client:       *client,
		Category:     *category,
		WorkSchedule: *workSchedule,
		Specialties:  *specialties,
	}, nil
}
