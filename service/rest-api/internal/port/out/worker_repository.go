package out

import (
	"context"
	"service/rest-api/internal/core/domain"

	"github.com/google/uuid"
)

type WorkerRepository interface {
	GetAllWorkers(ctx context.Context) (*[]domain.Worker, error)
	GetWorkerById(ctx context.Context, id *uuid.UUID) (*domain.Worker, error)
	GetWorkScheduleById(ctx context.Context, id *uuid.UUID) (*domain.WorkSchedule, error)
	GetSpecialitiesByClientId(ctx context.Context, id *uuid.UUID) (*[]domain.Specialty, error)
	CreateWorker(ctx context.Context, user domain.WorkerRequest) (*uuid.UUID, error)
}
