package repository

import (
	"log"
	"service/rest-api/internal/core/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ClientRepository struct {
	db Database
}

func NewClientRepository(database Database) ClientRepository {
	return ClientRepository{
		db: database,
	}
}

func (r ClientRepository) GetAllClients() (*[]domain.Client, error) {
	query := `
		SELECT * FROM client
	`
	rows, err := r.db.query(query)
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

func (r ClientRepository) GetClientById(id *uuid.UUID) (*domain.Client, error) {
	query := `
		SELECT * FROM client WHERE id = @clientId
	`
	args := pgx.NamedArgs{
		"clientId": *id,
	}

	var client domain.Client
	err := r.db.queryRow(query, args).Scan(&client.Id, &client.Name, &client.EmailAddress, &client.CategoryId, &client.WorkScheduleId)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (r ClientRepository) GetCategoryById(id *uuid.UUID) (*domain.Category, error) {
	query := `
		SELECT * FROM category WHERE id = @id
	`
	args := pgx.NamedArgs{
		"id": *id,
	}

	var category domain.Category
	err := r.db.queryRow(query, args).Scan(&category.Id, &category.Rol)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r ClientRepository) GetWorkScheduleById(id *uuid.UUID) (*domain.WorkSchedule, error) {
	query := `
		SELECT * FROM workSchedule WHERE id = @id
	`
	args := pgx.NamedArgs{
		"id": *id,
	}

	var schedule domain.WorkSchedule
	err := r.db.queryRow(query, args).Scan(&schedule.Id, &schedule.StartTime, &schedule.EndTime)
	if err != nil {
		return nil, err
	}
	return &schedule, nil
}

// TODO: Test when row is empty
func (r ClientRepository) GetSpecialitiesByClientId(id *uuid.UUID) (*[]domain.Specialty, error) {
	query := `
		SELECT * FROM specialty WHERE userId = @id
	`
	args := pgx.NamedArgs{
		"id": *id,
	}

	rows, err := r.db.query(query, args)
	if err != nil {
		log.Printf("Error, could not fetch data: %v", err)
		return nil, err
	}
	defer rows.Close()

	array, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Specialty])
	if err != nil {
		log.Printf("Error, could not create array: %v", err)
		return nil, err
	}

	return &array, nil
}
