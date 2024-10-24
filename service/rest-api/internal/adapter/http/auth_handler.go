package adapter

import (
	"fmt"
	"net/http"

	"service/rest-api/internal/core/domain"
	"service/rest-api/internal/port/in"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	service       in.AuthService
	cognitoClient CognitoClient
}

func NewAuthHandler(s in.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) SignIn(c echo.Context) error {
	var user domain.UserLogin

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	token, err := h.service.SignIn(user)
	if err != nil {
		return c.JSON(
			http.StatusUnauthorized,
			map[string]string{
				"error": "token:" + token + fmt.Sprintln(err) + h.cognitoClient.appClientID,
			})
	}

	return c.JSON(http.StatusOK, map[string]string{"access_token": token})
}
