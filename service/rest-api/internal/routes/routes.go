package routes

import (
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

	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := adapter.NewUserHandler(userService)

	e.GET("/user", userHandler.GetAllUsers)
}
