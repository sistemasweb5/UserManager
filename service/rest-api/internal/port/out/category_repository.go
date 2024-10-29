package out

import (
	"context"
	"service/rest-api/internal/core/domain"
)

type CategoryRepository interface {
	InsertCategory(ctx context.Context, category *domain.Category) error
	GetAllCategory(ctx context.Context) (*[]domain.Category, error)
}
