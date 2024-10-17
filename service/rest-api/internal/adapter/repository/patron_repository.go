package repository

import (
	"log"

	"service/rest-api/internal/core/domain"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db Database
}

func NewUserRepository(database Database) UserRepository {
	return UserRepository{
		db: database,
	}
}

func (r UserRepository) GetAll() (*[]domain.User, error) {
	query := `
		SELECT * FROM user
	`
	rows, err := r.db.query(query)
	if err != nil {
		log.Printf("Error, could not fetch data: %v", err)
		return nil, err
	}
	defer rows.Close()

	array, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.User])
	if err != nil {
		log.Printf("Error, could not create array: %v", err)
		return nil, err
	}
	
	return &array, nil
}
