package out

import (
	"service/rest-api/internal/core/domain"

	"github.com/google/uuid"
)

type ClientRepository interface {
	GetAllClients() (*[]domain.Client, error)
	GetClientById(id *uuid.UUID) (*domain.Client, error)
	GetCategoryById(id *uuid.UUID) (*domain.Category, error)
	GetWorkScheduleById(id *uuid.UUID) (*domain.WorkSchedule, error)
	GetSpecialitiesByClientId(id *uuid.UUID) (*[]domain.Specialty, error)
}
