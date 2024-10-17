package out

import (
	"service/rest-api/internal/core/domain"
)

type UserRepository interface {
	GetAll() (*[]domain.User, error)
}
