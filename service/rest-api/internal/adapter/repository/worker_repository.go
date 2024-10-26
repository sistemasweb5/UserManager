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

type WorkerRepository struct {
	db               *pgxpool.Pool
	workerCategoryId uuid.UUID
}

func NewWorkerRepository(conn *pgxpool.Pool, categoryId uuid.UUID) out.WorkerRepository {
	return &WorkerRepository{
		db:               conn,
		workerCategoryId: categoryId,
	}
}

func (w *WorkerRepository) GetSpecialitiesByClientId(ctx context.Context, id *uuid.UUID) (*[]domain.Specialty, error) {
	query := `
		SELECT * FROM specialty WHERE clientId = $1
	`
	userId := *id
	rows, err := w.db.Query(ctx, query, userId)
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

func (w *WorkerRepository) GetWorkScheduleById(ctx context.Context, id *uuid.UUID) (*domain.WorkSchedule, error) {
	query := `
		SELECT * FROM workSchedule WHERE id = @id
	`
	args := pgx.NamedArgs{
		"id": *id,
	}

	var schedule domain.WorkSchedule
	err := w.db.QueryRow(ctx, query, args).Scan(&schedule.Id, &schedule.StartTime, &schedule.EndTime)
	if err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (w *WorkerRepository) GetWorkerById(ctx context.Context, id *uuid.UUID) (*domain.Worker, error) {
	query := `
		SELECT client.id, client.name, client.emailAddress, client.workScheduleId
		FROM client
		WHERE client.categoryId = @categoryId AND client.id = @id
	`
	args := pgx.NamedArgs{
		"id":         *id,
		"categoryId": w.workerCategoryId,
	}

	var worker domain.Worker
	err := w.db.QueryRow(ctx, query, args).Scan(&worker.Id, &worker.Name, &worker.EmailAddress, &worker.WorkScheduleId)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}

	return &worker, nil
}

func (w *WorkerRepository) GetAllWorkers(ctx context.Context) (*[]domain.Worker, error) {
	query := `
		SELECT client.id, client.name, client.emailAddress, client.workScheduleId
		FROM client
		WHERE client.categoryId = $1
	`
	rows, err := w.db.Query(ctx, query, w.workerCategoryId)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	defer rows.Close()

	array, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Worker])
	if err != nil {
		log.Printf("Error, could not create array: %v", err)
		return nil, err
	}

	return &array, nil
}

func (w *WorkerRepository) CreateWorker(ctx context.Context, user domain.WorkerRequest) (*uuid.UUID, error) {
	id := uuid.New()

	query := `
		INSERT INTO client (id, name, emailAddress, categoryId, workScheduleId)
		VALUES (@id, @name, @emailAddress, @categoryId, @workScheduleId)
	`
	args := pgx.NamedArgs{
		"id":             id,
		"name":           user.Name,
		"emailAddress":   user.EmailAddress,
		"categoryId":     w.workerCategoryId,
		"workScheduleId": user.WorkScheduleId,
	}

	if _, err := w.db.Exec(ctx, query, args); err != nil {
		log.Printf("Error, could not insert data: %v", err)
		return nil, err
	}

	return &id, nil
}
