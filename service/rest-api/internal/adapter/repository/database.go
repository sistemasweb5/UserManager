package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	Connection *pgx.Conn
}

func (db *Database) execute(query string, args pgx.NamedArgs) error {
	if _, err := db.Connection.Exec(context.Background(), query, args); err != nil {
		log.Printf("Could not execute query: %v", err)
		return err
	}

	return nil
}

func (db *Database) query(query string) (pgx.Rows, error) {
	return db.Connection.Query(context.Background(), query)
}

func (db *Database) queryRow(query string, args pgx.NamedArgs) pgx.Row {
	return db.Connection.QueryRow(context.Background(), query, args)
}
