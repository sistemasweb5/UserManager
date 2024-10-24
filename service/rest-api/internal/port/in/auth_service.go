package in

import "service/rest-api/internal/core/domain"

type AuthService interface {
	SignIn(user domain.UserLogin) (string, error)
	Logout(accessToken string) error
}
