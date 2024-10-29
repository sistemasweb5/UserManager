package in

import (
	"context"
	"service/rest-api/internal/core/domain"

	"github.com/google/uuid"
)

// Deprecated: Use CategoryService, WorkerService and ApplicantService instead.
type ClientService interface {
	GetById(ctx context.Context, id *uuid.UUID) (*domain.ClientResponse, error)
}
