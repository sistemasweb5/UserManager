package service

import (
	"context"
	"service/rest-api/internal/core/domain"
	"service/rest-api/internal/port/in"
	"service/rest-api/internal/port/out"
)

type ApplicantService struct {
	repo out.ApplicantRepository
}

func NewApplicantService(repo out.ApplicantRepository) in.ApplicantService {
	return &ApplicantService{repo: repo}
}

func (a *ApplicantService) GetAll(ctx context.Context) (*[]domain.ApplicantResponse, error) {
	return a.repo.GetAllApplicants(ctx)
}
