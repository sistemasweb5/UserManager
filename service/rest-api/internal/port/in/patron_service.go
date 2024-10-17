package in

import (
	"service/rest-api/internal/core/domain"
)

type UserService interface {
	GetAll() (*[]domain.User, error)
}
