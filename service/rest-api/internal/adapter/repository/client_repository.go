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

func (r ClientRepository) GetAll() (*[]domain.Client, error) {
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

func (r ClientRepository) GetById(id *uuid.UUID) (*domain.Client, error) {
	query := `
		SELECT * FROM client WHERE id = @clientId
	`
	args := pgx.NamedArgs{
        "clientId": *id,
    }

	var client domain.Client
	err := r.db.queryRow(query, args).Scan(&client.Id, &client.Name, &client.EmailAddress, &client.CategoryId)
	if err != nil {
		return nil, err
	}
	return &client, nil
}
