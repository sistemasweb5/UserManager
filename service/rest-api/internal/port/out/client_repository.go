package out

import (
	"service/rest-api/internal/core/domain"
)

type ClientRepository interface {
	GetAll() (*[]domain.Client, error)
}
