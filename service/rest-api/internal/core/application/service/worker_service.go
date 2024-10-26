package service

import (
	"context"
	"service/rest-api/internal/core/domain"
	"service/rest-api/internal/port/in"
	"service/rest-api/internal/port/out"

	"github.com/google/uuid"
)

type WorkerService struct {
	workerRepo out.WorkerRepository
}

func NewWorkerService(repo out.WorkerRepository) in.WorkerService {
	return WorkerService{
		workerRepo: repo,
	}
}

func (w WorkerService) GetAllWorker(ctx context.Context) (*[]domain.WorkerResponse, error) {
	workers, err := w.workerRepo.GetAllWorkers(ctx)
	if err != nil {
		return nil, err
	}

	var workerResponses []domain.WorkerResponse
	for _, worker := range *workers {
		workerResponse, err := w.GetWorkerById(ctx, &worker.Id)
		if err != nil {
			return nil, err
		}
		workerResponses = append(workerResponses, *workerResponse)
	}

	return &workerResponses, nil
}

func (w WorkerService) GetWorkerById(ctx context.Context, id *uuid.UUID) (*domain.WorkerResponse, error) {
	worker, err := w.workerRepo.GetWorkerById(ctx, id)
	if err != nil {
		return nil, err
	}

	schedule, err := w.workerRepo.GetWorkScheduleById(ctx, &worker.WorkScheduleId)
	if err != nil {
		return nil, err
	}

	specialties, err := w.workerRepo.GetSpecialitiesByClientId(ctx, &worker.Id)
	if err != nil {
		return nil, err
	}

	return &domain.WorkerResponse{
		Id:           worker.Id,
		Name:         worker.Name,
		EmailAddress: worker.EmailAddress,
		WorkSchedule: *schedule,
		Specialties:  *specialties,
	}, nil
}

func (w WorkerService) CreateWorker(ctx context.Context, request domain.WorkerRequest) (*uuid.UUID, error) {
	return w.workerRepo.CreateWorker(ctx, request)
}
