package adapter

import (
	"log"
	"net/http"
	"service/rest-api/internal/port/in"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ClientHandler struct {
	service in.ClientService
}

func NewClientHandler(s in.ClientService) ClientHandler {
	return ClientHandler{service: s}
}

func (handler *ClientHandler) GetAllClients(context echo.Context) error {
	clients, err := handler.service.GetAll(context.Request().Context())
	if err != nil {
		log.Printf("Error: %v", err)

		return context.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Unable to retrieve clients",
		})
	}
	return context.JSON(http.StatusOK, clients)
}

func (handler *ClientHandler) GetClientById(context echo.Context) error {
	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		log.Printf("Could not parse id: %v", err)
		return err
	}

	client, err := handler.service.GetById(context.Request().Context(), &id, )
	if err != nil {
		log.Printf("Could not get client: %v", err)
		return err
	}

	return context.JSON(http.StatusOK, client)
}
