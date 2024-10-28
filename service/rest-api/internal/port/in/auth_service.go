package in

import "service/rest-api/internal/core/domain"

type AuthService interface {
	SignIn(user domain.UserLogin) (string, error)
	Logout(accessToken string) error
	GetUserIdByToken(accessToken string) (string, error)
	SignUp(user domain.UserSignUp) (string, error)
	ConfirmAccount(user domain.UserConfirmation) error
}
