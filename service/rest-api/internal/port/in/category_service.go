package in

import (
	"context"
	"service/rest-api/internal/core/domain"

	"github.com/google/uuid"
)

type CategoryService interface {
	GetAllCategory(ctx context.Context) (*[]domain.Category, error)
	IsColdStart(ctx context.Context) (bool, error)
	Populate(ctx context.Context) (*uuid.UUID, *uuid.UUID, error)
	Realize(ctx context.Context) (*uuid.UUID, *uuid.UUID, error)
}
