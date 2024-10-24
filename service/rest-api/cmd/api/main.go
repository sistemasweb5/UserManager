package main

import (
	"context"
	"log"
	"os"
	"service/rest-api/internal/routes"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

func main() {
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

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	e := echo.New()
	routes.RegisterRoutes(e, pgxConnPool)
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Could not start api: %v", e)
	}
}
