package out

import (
	"context"
	"service/rest-api/internal/core/domain"

	"github.com/google/uuid"
)

type ApplicantRepository interface {
	GetAllApplicants(ctx context.Context) (*[]domain.Applicant, error)
	GetApplicantById(ctx context.Context, id *uuid.UUID) (*domain.Applicant, error)
	CreateApplicant(ctx context.Context, user domain.ApplicantRequest) (*uuid.UUID, error)
}
