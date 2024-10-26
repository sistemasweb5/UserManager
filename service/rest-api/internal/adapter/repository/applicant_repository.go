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

type ApplicantRepository struct {
	db                  *pgxpool.Pool
	applicantCategoryId *uuid.UUID
}

func NewApplicantRepository(conn *pgxpool.Pool, categoryId *uuid.UUID) out.ApplicantRepository {
	return ApplicantRepository{
		db:                  conn,
		applicantCategoryId: categoryId,
	}
}

func (r ApplicantRepository) GetAllApplicants(ctx context.Context) (*[]domain.Applicant, error) {
	query := `
		SELECT client.id, client.name, client.emailAddress
		FROM client
		WHERE client.categoryId = $1
	`
	rows, err := r.db.Query(ctx, query, r.applicantCategoryId)
	if err != nil {
		log.Printf("Error, could not fetch data: %v", err)
		return nil, err
	}
	defer rows.Close()

	array, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Applicant])
	if err != nil {
		log.Printf("Error, could not create array: %v", err)
		return nil, err
	}

	return &array, nil
}

func (r ApplicantRepository) GetApplicantById(ctx context.Context, id *uuid.UUID) (*domain.Applicant, error) {
	query := `
		SELECT client.id, client.name, client.emailAddress
		FROM client
		WHERE client.categoryId = @categoryId AND client.id = @id
	`
	args := pgx.NamedArgs{
		"id":         *id,
		"categoryId": *r.applicantCategoryId,
	}

	var user domain.Applicant
	err := r.db.QueryRow(ctx, query, args).Scan(&user.Id, &user.Name, &user.EmailAddress)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}

	return &user, nil
}

func (r ApplicantRepository) CreateApplicant(ctx context.Context, user domain.ApplicantRequest) (*uuid.UUID, error) {
	id := uuid.New()

	query := `
		INSERT INTO client (id, name, emailAddress, categoryId)
		VALUES (@id, @name, @emailAddress, @categoryId)
	`
	args := pgx.NamedArgs{
		"id":             id,
		"name":           user.Name,
		"emailAddress":   user.EmailAddress,
		"categoryId":     r.applicantCategoryId,
	}

	if _, err := r.db.Exec(ctx, query, args); err != nil {
		log.Printf("Error, could not insert data: %v", err)
		return nil, err
	}

	return &id, nil
}

