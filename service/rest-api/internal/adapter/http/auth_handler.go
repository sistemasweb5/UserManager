package adapter

import (
	"fmt"
	"net/http"
	"strings"

	"service/rest-api/internal/core/domain"
	"service/rest-api/internal/port/in"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	service in.AuthService
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
				"error": "token:" + token + fmt.Sprintln(err),
			})
	}

	return c.JSON(http.StatusOK, map[string]string{"access_token": token})
}

func (h *AuthHandler) Logout(c echo.Context) error {
	accessToken := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")
	if accessToken == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Authorization token required"})
	}

	err := h.service.Logout(accessToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": fmt.Sprintln(err)})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Successfully logged out"})
}
