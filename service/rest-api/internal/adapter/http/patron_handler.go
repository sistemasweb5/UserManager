package adapter

import (
	"log"
	"net/http"
	"service/rest-api/internal/port/in"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service in.UserService
}

func NewUserHandler(s in.UserService) UserHandler {
	return UserHandler{service: s}
}

func (handler *UserHandler) GetAllUsers(context echo.Context) error {
	users, err := handler.service.GetAll()
	if err != nil {
		log.Printf("Could not get all users: %v", err)

		return context.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Unable to retrieve users",
		})
	}
	return context.JSON(http.StatusOK, users)
}
