package out

import (
	"service/rest-api/internal/core/domain"

	"github.com/google/uuid"
)

type ClientRepository interface {
	GetAll() (*[]domain.Client, error)
	GetById(id *uuid.UUID) (*domain.Client, error)
	Create(client *domain.Client) error
}
