package main

import (
	"context"
	"log"
	"os"
	"service/rest-api/internal/adapter/repository"
	"service/rest-api/internal/core/domain"
	"service/rest-api/internal/routes"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

func populateCategories(ctx context.Context, db *pgxpool.Pool) (uuid.UUID, uuid.UUID, error) {
	categoryRepo := repository.NewCategoryRepository(db)

	workerCategoryId := uuid.New()
	applicantCategoryId := uuid.New()

	err := categoryRepo.InsertCategory(ctx, &domain.Category{
		Id:  applicantCategoryId,
		Rol: "applicant",
	})
	if err != nil {
		log.Fatalf("ERROR: Could not initialize database: %v", err)
	}

	err = categoryRepo.InsertCategory(ctx, &domain.Category{
		Id:  workerCategoryId,
		Rol: "worker",
	})
	if err != nil {
		log.Fatalf("ERROR: Could not initialize database: %v", err)
	}

	return workerCategoryId, applicantCategoryId, nil
}

func setupDB() *pgxpool.Pool {
	pgxConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}


	pgxConnPool, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	return pgxConnPool
}

func main() {
	pgxConnPool := setupDB()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	routes.RegisterRoutes(e, pgxConnPool)
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Could not start api: %v", e)
	}
}
