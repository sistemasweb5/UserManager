package service

import (
	"context"
	"fmt"
	"log"
	"service/rest-api/internal/core/domain"
	"service/rest-api/internal/port/in"
	"service/rest-api/internal/port/out"

	"github.com/google/uuid"
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

func (c CategoryService) IsColdStart(ctx context.Context) (bool, error) {
	categories, err := c.GetAllCategory(ctx)
	if err != nil {
		log.Print(err)
		return false, err
	}
	return (len(*categories) != 2), err
}

func (c CategoryService) Realize(ctx context.Context) (*uuid.UUID, *uuid.UUID, error) {
	categories, err := c.GetAllCategory(ctx)
	if err != nil {
		return nil, nil, err
	}

	var workerCategory domain.Category
	var applicantCategory domain.Category
	workerCategory = (*categories)[0]
	if workerCategory.Rol != "worker" {
		workerCategory = (*categories)[1]
		applicantCategory = (*categories)[0]
	}
	if workerCategory.Rol != "worker" {
		return nil, nil, fmt.Errorf("Found a different category.\n %v", *categories)
	}
	return &applicantCategory.Id, &workerCategory.Id, nil
}

func (c CategoryService) Populate(ctx context.Context) (*uuid.UUID, *uuid.UUID, error) {
	applicantCategoryId := uuid.New()
	workerCategoryId := uuid.New()
	err := c.categoryRepo.InsertCategory(ctx, &domain.Category{
		Id:  applicantCategoryId,
		Rol: "applicant",
	})
	if err != nil {
		log.Fatalf("ERROR: Could not initialize database: %v", err)
	}
	err = c.categoryRepo.InsertCategory(ctx, &domain.Category{
		Id:  workerCategoryId,
		Rol: "worker",
	})
	if err != nil {
		log.Fatalf("ERROR: Could not initialize database: %v", err)
	}

	return &applicantCategoryId, &workerCategoryId, nil
}
