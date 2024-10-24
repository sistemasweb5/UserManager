package repository

import (
	"context"
	"log"
	"service/rest-api/internal/core/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ApplicantRepository struct {
	db *pgxpool.Pool
}

func NewApplicantRepository(conn *pgxpool.Pool) *ApplicantRepository {
	return &ApplicantRepository{
		db: conn,
	}
}

func (r *ApplicantRepository) GetAllApplicants(ctx context.Context) (*[]domain.ApplicantResponse, error) {
	query := `
		SELECT client.id, client.name, client.emailAddress FROM client
		INNER JOIN category ON client.categoryId = category.id
		WHERE category.rol = 'applicant'
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error, could not fetch data: %v", err)
		return nil, err
	}
	defer rows.Close()

	array, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.ApplicantResponse])
	if err != nil {
		log.Printf("Error, could not create array: %v", err)
		return nil, err
	}

	return &array, nil
}
