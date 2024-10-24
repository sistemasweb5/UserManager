package adapter

import (
	"log"
	"net/http"
	"service/rest-api/internal/core/domain"
	"service/rest-api/internal/port/in"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	service in.AuthService
}

func NewAuthHandler(s in.AuthService) AuthHandler {
	return AuthHandler{service: s}
}

func (h *AuthHandler) SignUp(c echo.Context) error {
	var user domain.User
	if err := c.Bind(&user); err != nil {
		log.Printf("Failed to bind user data: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid user data",
		})
	}

	if err := h.service.SignUp(&user); err != nil {
		log.Printf("Failed to sign up user: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "User created successfully",
	})
}

func (h *AuthHandler) ConfirmAccount(c echo.Context) error {
	var confirmation domain.UserConfirmation
	if err := c.Bind(&confirmation); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid confirmation data",
		})
	}

	if err := h.service.ConfirmAccount(&confirmation); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to confirm account",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Account confirmed successfully",
	})
}
