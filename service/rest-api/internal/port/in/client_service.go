package in

import (
	"context"
	"service/rest-api/internal/core/domain"

	"github.com/google/uuid"
)

type ClientService interface {
	GetAll(ctx context.Context) (*[]domain.ClientResponse, error)
	GetById(ctx context.Context, id *uuid.UUID) (*domain.ClientResponse, error)
}
