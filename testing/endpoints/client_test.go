package endpoints_test

import (
	"api/testing/endpoints"
	"context"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func databaseSetup() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("TEST_DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	db := endpoints.Database{
		Connection: conn,
	}

	categoryId := uuid.NewString()
	userId := uuid.NewString()
	workScheduleId := uuid.NewString()
	client := endpoints.Applicant{
		Id:           userId,
		Name:         "Slim shady",
		EmailAddress: "fake_email@mail.com",
		CategoryId:   categoryId,
	}

	schedule := endpoints.WorkSchedule{
		Id:        workScheduleId,
		StartTime: "9.00",
		EndTime:   "17.00",
	}
	specialty := endpoints.Specialty{
		Id:       uuid.NewString(),
		Name:     "Plumber",
		ClientId: userId,
	}

	if err := db.InsertWorkSchedule(&schedule); err != nil {
		log.Fatalf("Could not populate table: %v", err)
	}
	if err := db.InsertApplicant(&client); err != nil {
		log.Fatalf("Could not populate table: %v", err)
	}
	if err := db.InsertSpecialty(&specialty); err != nil {
		log.Fatalf("Could not populate table: %v", err)
	}
}

func TestMain(m *testing.M) {
	log.SetFlags(log.Lshortfile)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseSetup()
	code := m.Run()
	os.Exit(code)
}
