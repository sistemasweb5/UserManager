package in

import (
	"context"
	"service/rest-api/internal/core/domain"
)

type CategoryService interface {
	GetAllCategory(ctx context.Context) (*[]domain.Category, error)
}
