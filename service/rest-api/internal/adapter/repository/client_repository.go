package repository

import (
	"log"

	"service/rest-api/internal/core/domain"
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
