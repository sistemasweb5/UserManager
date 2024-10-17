package adapter

import (
	"log"
	"net/http"
	"service/rest-api/internal/port/in"

	"github.com/labstack/echo/v4"
)

type ClientHandler struct {
	service in.ClientService
}

func NewClientHandler(s in.ClientService) ClientHandler {
	return ClientHandler{service: s}
}

func (handler *ClientHandler) GetAllClients(context echo.Context) error {
	clients, err := handler.service.GetAll()
	if err != nil {
		log.Printf("Could not get all clients: %v", err)

		return context.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Unable to retrieve clients",
		})
	}
	return context.JSON(http.StatusOK, clients)
}
