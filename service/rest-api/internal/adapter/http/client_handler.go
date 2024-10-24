package adapter

import (
	"log"
	"net/http"
	"service/rest-api/internal/core/domain"
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
	clients, err := handler.service.GetAll()
	if err != nil {
		log.Printf("Could not get all clients: %v", err)

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

	client, err := handler.service.GetById(&id)
	if err != nil {
		log.Printf("Could not get client: %v", err)
		return err
	}

	return context.JSON(http.StatusOK, client)
}

func (handler *ClientHandler) CreateUser(context echo.Context) error {
	var client domain.Client
	if err := context.Bind(&client); err != nil {
		log.Printf("Failed to bind user data: %v", err)
		log.Printf("Request body: %v", context.Request().Body)
		return context.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid user data",
		})
	}

	log.Printf("Client received: %+v", client)

	if client.Name == "" || client.EmailAddress == "" || client.CategoryId == uuid.Nil {
		return context.JSON(http.StatusBadRequest, map[string]string{
			"message": "All fields are required",
		})
	}

	if err := handler.service.Create(&client); err != nil {
		log.Printf("Failed to save user: %v", err)
		return context.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to create user",
		})
	}

	return context.JSON(http.StatusCreated, map[string]interface{}{
		"message": "User created successfully",
	})
}
