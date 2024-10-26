package service

import (
	"context"
	"service/rest-api/internal/core/domain"
	"service/rest-api/internal/port/in"
	"service/rest-api/internal/port/out"
)

type CategoryService struct {
	categoryRepo out.CategoryRepository
}

func NewCategoryService(category out.CategoryRepository) in.CategoryService {
	return CategoryService{
		categoryRepo: category,
	}
}

func (c CategoryService) GetAllCategory(ctx context.Context) (*[]domain.Category, error) {
	return c.categoryRepo.GetAllCategory(ctx)
}
