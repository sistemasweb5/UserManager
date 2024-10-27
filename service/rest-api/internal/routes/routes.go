package routes

import (
	"context"
	"log"
	"os"
	adapter "service/rest-api/internal/adapter/http"
	"service/rest-api/internal/adapter/repository"
	"service/rest-api/internal/core/application/service"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, conn *pgxpool.Pool) {
	clientHandler, err := setupClientService(context.TODO(), conn)
	if err != nil {
		log.Fatalf("Could not startup: %v", err)
	}

	cognitoClient := repository.NewCognitoClient(os.Getenv("COGNITO_APP_CLIENT_ID"))
	authService := service.NewAuthService(cognitoClient)
	authHandler := adapter.NewAuthHandler(authService)

	e.POST("/user/login", authHandler.SignIn)
	e.GET("/user/id", authHandler.GetUserIdByToken)
	e.POST("/user/logout", authHandler.Logout)
	e.POST("/user/signup", authHandler.SignUp)
	e.POST("/user/confirm", authHandler.ConfirmAccount)

	e.GET("/client/:id", clientHandler.GetClientById)
	e.GET("/client/category", clientHandler.GetAllCategory)
	e.GET("/client/applicant", clientHandler.GetAllApplicant)
	e.GET("/client/applicant/:id", clientHandler.GetApplicantById)
	e.GET("/client/worker", clientHandler.GetAllWorker)
	e.GET("/client/worker/:id", clientHandler.GetWorkerById)

	e.POST("/client/worker", clientHandler.CreateWorker)
	e.POST("/client/applicant", clientHandler.CreateApplicant)
}

func setupClientService(ctx context.Context, conn *pgxpool.Pool) (adapter.ClientHandler, error) {
	categoryRepo := repository.NewCategoryRepository(conn)
	categoryService := service.NewCategoryService(categoryRepo)
	isColdStart, err := categoryService.IsColdStart(ctx)
	if err != nil {
		return adapter.ClientHandler{}, err
	}
	var workerCategoryId *uuid.UUID
	var applicantCategoryId *uuid.UUID
	if isColdStart {
		applicantCategoryId, workerCategoryId, err = categoryService.Populate(ctx)
		if err != nil {
			return adapter.ClientHandler{}, err
		}
	} else {
		applicantCategoryId, workerCategoryId, err = categoryService.Realize(ctx)
		if err != nil {
			return adapter.ClientHandler{}, err
		}
	}

	applicantRepo := repository.NewApplicantRepository(conn, applicantCategoryId)
	applicantService := service.NewApplicantService(applicantRepo)
	workerRepo := repository.NewWorkerRepository(conn, workerCategoryId)
	workerService := service.NewWorkerService(workerRepo)

	clienRepo := repository.NewClientRepository(conn)
	clientService := service.NewClientService(clienRepo)
	return adapter.NewClientHandler(applicantService, workerService, categoryService, clientService), nil
}
