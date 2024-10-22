package routes

import (
	adapter "service/rest-api/internal/adapter/http"
	"service/rest-api/internal/adapter/repository"
	"service/rest-api/internal/core/application/service"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, conn *pgxpool.Pool) {
	clientRepo := repository.NewClientRepository(conn)
	clientService := service.NewClientService(clientRepo)
	clientHandler := adapter.NewClientHandler(clientService)

	e.GET("/client", clientHandler.GetAllClients)
	e.GET("/client/:id", clientHandler.GetClientById)
}
