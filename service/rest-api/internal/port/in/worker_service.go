package in

import (
	"context"
	"service/rest-api/internal/core/domain"

	"github.com/google/uuid"
)

type WorkerService interface {
	GetAllWorker(ctx context.Context) (*[]domain.WorkerResponse, error)
	GetWorkerById(ctx context.Context, id *uuid.UUID) (*domain.WorkerResponse, error)
	CreateWorker(ctx context.Context, request domain.WorkerRequest) (*uuid.UUID, error)
}
