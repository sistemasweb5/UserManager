package in

import (
	"service/rest-api/internal/core/domain"

	"github.com/google/uuid"
)

type ClientService interface {
	GetAll() (*[]domain.ClientResponse, error)
	GetById(id *uuid.UUID) (*domain.ClientResponse, error)
}
