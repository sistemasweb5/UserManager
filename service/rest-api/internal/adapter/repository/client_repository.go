package repository

import (
	"context"
	"log"
	"service/rest-api/internal/core/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ClientRepository struct {
	db *pgxpool.Pool
}

func NewClientRepository(conn *pgxpool.Pool) ClientRepository {
	return ClientRepository{
		db: conn,
	}
}

func (r ClientRepository) GetAllClients(ctx context.Context) (*[]domain.Client, error) {
	query := `
		SELECT * FROM client
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error, could not fetch data: %v", err)
		return nil, err
	}
	defer rows.Close()

	array, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Client])
	if err != nil {
		log.Printf("Error, could not create array: %v", err)
		return nil, err
	}

	return &array, nil
}

func (r ClientRepository) GetClientById(ctx context.Context, id *uuid.UUID) (*domain.Client, error) {
	query := `
		SELECT * FROM client WHERE id = @clientId
	`
	args := pgx.NamedArgs{
		"clientId": *id,
	}

	var client domain.Client
	err := r.db.QueryRow(ctx, query, args).Scan(&client.Id, &client.Name, &client.EmailAddress, &client.CategoryId, &client.WorkScheduleId)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return &client, nil
}

func (r ClientRepository) GetCategoryById(ctx context.Context, id *uuid.UUID) (*domain.Category, error) {
	query := `
		SELECT * FROM category WHERE id = @id
	`
	args := pgx.NamedArgs{
		"id": *id,
	}

	var category domain.Category
	err := r.db.QueryRow(ctx, query, args).Scan(&category.Id, &category.Rol)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r ClientRepository) GetWorkScheduleById(ctx context.Context, id *uuid.UUID) (*domain.WorkSchedule, error) {
	query := `
		SELECT * FROM workSchedule WHERE id = @id
	`
	args := pgx.NamedArgs{
		"id": *id,
	}

	var schedule domain.WorkSchedule
	err := r.db.QueryRow(ctx, query, args).Scan(&schedule.Id, &schedule.StartTime, &schedule.EndTime)
	if err != nil {
		return nil, err
	}
	return &schedule, nil
}

// TODO: Test when row is empty
func (r ClientRepository) GetSpecialitiesByClientId(ctx context.Context, id *uuid.UUID) (*[]domain.Specialty, error) {
	query := `
		SELECT * FROM specialty WHERE clientId = $1
	`
	userId := *id
	rows, err := r.db.Query(ctx, query, userId)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	defer rows.Close()

	array, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Specialty])
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}

	return &array, nil
}
