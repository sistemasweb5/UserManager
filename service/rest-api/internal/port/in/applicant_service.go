package in

import (
	"context"
	"service/rest-api/internal/core/domain"

	"github.com/google/uuid"
)

type ApplicantService interface {
	GetApplicantById(ctx context.Context, id *uuid.UUID) (*domain.Applicant, error)
	GetAllApplicant(ctx context.Context) (*[]domain.Applicant, error)
	CreateApplicant(ctx context.Context, applicant domain.ApplicantRequest) (*uuid.UUID, error)
}
