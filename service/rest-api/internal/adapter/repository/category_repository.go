package repository

import (
	"context"
	"log"
	"service/rest-api/internal/core/domain"
	"service/rest-api/internal/port/out"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryRepository struct {
	db           *pgxpool.Pool
	workerId     uuid.UUID
	appplicantId uuid.UUID
}

func NewCategoryRepository(conn *pgxpool.Pool) out.CategoryRepository {
	return CategoryRepository{
		db: conn,
	}
}

func (r CategoryRepository) GetAllCategory(ctx context.Context) (*[]domain.Category, error) {
	query := `
		SELECT * FROM category
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error, could not fetch data: %v", err)
		return nil, err
	}
	defer rows.Close()

	array, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Category])
	if err != nil {
		log.Printf("Error, could not create array: %v", err)
		return nil, err
	}

	return &array, nil
}

func (r CategoryRepository) InsertCategory(ctx context.Context, category *domain.Category) error {
	query := `
		INSERT INTO category (id, rol) VALUES (@id, @rol)
	`
	args := pgx.NamedArgs{
		"id":  category.Id,
		"rol": category.Rol,
	}
	if _, err := r.db.Exec(ctx, query, args); err != nil {
		return err
	}
	return nil
}
