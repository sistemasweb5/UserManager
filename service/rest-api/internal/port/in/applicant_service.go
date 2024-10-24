package in

import (
	"context"
	"service/rest-api/internal/core/domain"
)

type ApplicantService interface {
	GetAll(ctx context.Context) (*[]domain.ApplicantResponse, error)
}
