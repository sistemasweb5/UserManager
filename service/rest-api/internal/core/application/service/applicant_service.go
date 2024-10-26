package service

import (
	"context"
	"github.com/google/uuid"
	"service/rest-api/internal/core/domain"
	"service/rest-api/internal/port/in"
	"service/rest-api/internal/port/out"
)

type ApplicantService struct {
	applicantRepo out.ApplicantRepository
}

func NewApplicantService(repo out.ApplicantRepository) in.ApplicantService {
	return ApplicantService{
		applicantRepo: repo,
	}
}

func (a ApplicantService) GetAllApplicant(ctx context.Context) (*[]domain.Applicant, error) {
	return a.applicantRepo.GetAllApplicants(ctx)
}

func (a ApplicantService) GetApplicantById(ctx context.Context, id *uuid.UUID) (*domain.Applicant, error) {
	return a.applicantRepo.GetApplicantById(ctx, id)
}

func (a ApplicantService) CreateApplicant(ctx context.Context, applicant domain.ApplicantRequest) (*uuid.UUID, error) {
	return a.applicantRepo.CreateApplicant(ctx, applicant)
}

