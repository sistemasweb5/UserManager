package endpoints

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	Connection *pgx.Conn
}

func NewDatabase() *Database {
	conn, err := pgx.Connect(context.Background(), os.Getenv("TEST_DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	return &Database{
		Connection: conn,
	}
}

func (db *Database) execute(query string, args pgx.NamedArgs) error {
	if _, err := db.Connection.Exec(context.Background(), query, args); err != nil {
		return err
	}

	return nil
}

func (db *Database) InsertClient(client *Client) error {
	query := `
        INSERT INTO client (id, name, emailAddress, categoryId, workScheduleId) VALUES (@id, @name, @emailAddress, @categoryId, @workScheduleId)
    `
	args := pgx.NamedArgs{
		"id":             client.Id,
		"name":           client.Name,
		"emailAddress":   client.EmailAddress,
		"categoryId":     client.CategoryId,
		"workScheduleId": client.WorkScheduleId,
	}

	return db.execute(query, args)
}

func (db *Database) InsertCategory(category *Category) error {
	query := `
        INSERT INTO category (id, rol) VALUES (@id, @rol)
    `
	args := pgx.NamedArgs{
		"id":  category.Id,
		"rol": category.Rol,
	}

	return db.execute(query, args)
}

func (db *Database) InsertWorkSchedule(schedule *WorkSchedule) error {
	query := `
        INSERT INTO workSchedule (id, startTime, endTime) VALUES (@id, @startTime, @endTime)
    `
	args := pgx.NamedArgs{
		"id":        schedule.Id,
		"startTime": schedule.StartTime,
		"endTime":   schedule.EndTime,
	}

	return db.execute(query, args)
}

func (db *Database) InsertSpecialty(specialty *Specialty) error {
	query := `
        INSERT INTO specialty (id, name, clientId) VALUES (@id, @name, @categoryId)
    `
	args := pgx.NamedArgs{
		"id":     specialty.Id,
		"name":   specialty.Name,
		"categoryId": specialty.ClientId,
	}

	return db.execute(query, args)
}
