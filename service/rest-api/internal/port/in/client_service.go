package in

import (
	"service/rest-api/internal/core/domain"
)

type ClientService interface {
	GetAll() (*[]domain.Client, error)
}
