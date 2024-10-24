package out

import "service/rest-api/internal/core/domain"

type AuthRepository interface {
	SignUp(user *domain.User) error
	ConfirmAccount(confirmation *domain.UserConfirmation) error
}
