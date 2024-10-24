package in

import "service/rest-api/internal/core/domain"

type AuthService interface {
	SignUp(user *domain.User) error
	ConfirmAccount(confirmation *domain.UserConfirmation) error
}
