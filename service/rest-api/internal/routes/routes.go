package routes

import (
	"os"
	adapter "service/rest-api/internal/adapter/http"
	"service/rest-api/internal/adapter/repository"
	"service/rest-api/internal/core/application/service"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, conn *pgx.Conn) {
	database := repository.Database{
		Connection: conn,
	}

	clientRepo := repository.NewClientRepository(database)
	clientService := service.NewClientService(clientRepo)
	clientHandler := adapter.NewClientHandler(clientService)

	cognitoClient := adapter.NewCognitoClient(os.Getenv("COGNITO_APP_CLIENT_ID"))
	authService := service.NewAuthService(cognitoClient)
	authHandler := adapter.NewAuthHandler(authService)

	e.POST("/user/login", authHandler.SignIn)
	e.GET("/client", clientHandler.GetAllClients)
	e.GET("/client/:id", clientHandler.GetClientById)
}
