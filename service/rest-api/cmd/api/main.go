package main

import (
	"context"
	"log"
	"os"

	"service/rest-api/internal/routes"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	e := echo.New()
	routes.RegisterRoutes(e, conn)
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Could not start api: %v", e)
	}
}
