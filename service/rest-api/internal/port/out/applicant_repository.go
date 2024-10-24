package out

import (
	"context"
	"service/rest-api/internal/core/domain"
)

type ApplicantRepository interface {
	GetAllApplicants(ctx context.Context) (*[]domain.ApplicantResponse, error)
}
