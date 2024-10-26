package routes

import (
	"context"
	"os"
	adapter "service/rest-api/internal/adapter/http"
	"service/rest-api/internal/adapter/repository"
	"service/rest-api/internal/core/application/service"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, conn *pgxpool.Pool) {
	clientHandler := setupClientService(context.TODO(), conn)

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

func setupClientService(ctx context.Context, conn *pgxpool.Pool) adapter.ClientHandler {
	applicantCategoryId, _ := uuid.Parse("03619b71-e334-4db9-a1e9-f450d3854b61")
	workerCategoryId, _ := uuid.Parse("367a2064-9b0a-4d77-ab29-1410cb8f84ec")

	categoryRepo := repository.NewCategoryRepository(conn)
	applicantRepo := repository.NewApplicantRepository(conn, &applicantCategoryId)
	workerRepo := repository.NewWorkerRepository(conn, workerCategoryId)
	clientRepo := repository.NewClientRepository(conn)

	// err := categoryRepo.InsertCategory(ctx, &domain.Category{
	// 	Id:  applicantCategoryId,
	// 	Rol: "applicant",
	// })
	// if err != nil {
	// 	log.Fatalf("ERROR: Could not initialize database: %v", err)
	// }
	// err = categoryRepo.InsertCategory(ctx, &domain.Category{
	// 	Id:  workerCategoryId,
	// 	Rol: "worker",
	// })
	// if err != nil {
	// 	log.Fatalf("ERROR: Could not initialize database: %v", err)
	// }

	workerService := service.NewWorkerService(workerRepo)
	applicantService := service.NewApplicantService(applicantRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	clientService := service.NewClientService(clientRepo)

	return adapter.NewClientHandler(applicantService, workerService, categoryService, clientService)
}
