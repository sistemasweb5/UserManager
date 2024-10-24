package out

import (
	"context"
	"service/rest-api/internal/core/domain"

	"github.com/google/uuid"
)

type ClientRepository interface {
	GetAllClients(ctx context.Context) (*[]domain.Client, error)
	GetClientById(ctx context.Context, id *uuid.UUID) (*domain.Client, error)
	GetCategoryById(ctx context.Context, id *uuid.UUID) (*domain.Category, error)
	GetWorkScheduleById(ctx context.Context, id *uuid.UUID) (*domain.WorkSchedule, error)
	GetSpecialitiesByClientId(ctx context.Context, id *uuid.UUID) (*[]domain.Specialty, error)
}
