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

type helper struct {
	db   *pgx.Conn
	myDb endpoints.Database
}

func (h *helper) makeWorkSchedule() (*string, error) {
	id := uuid.NewString()
	schedule := endpoints.WorkSchedule{
		Id:        id,
		StartTime: "9.00",
		EndTime:   "17.00",
	}
	return &id, h.myDb.InsertWorkSchedule(&schedule)
}

func (h *helper) makeSpecialty(userId string) (*string, error) {
	id := uuid.NewString()
	specialty := endpoints.Specialty{
		Id:       id,
		Name:     "Plumber",
		ClientId: userId,
	}

	return &id, h.myDb.InsertSpecialty(&specialty)
}

func (h *helper) fetchCategoryIds() (*string, *string, error) {
	query := `
		SELECT id FROM category WHERE rol = 'applicant'
	`
	var applicantId string
	err := h.db.QueryRow(context.Background(), query).Scan(&applicantId)
	if err != nil {
		return nil, nil, err
	}

	query = `
		SELECT id FROM category WHERE rol = 'worker'
	`
	var workerId string
	err = h.db.QueryRow(context.Background(), query).Scan(&workerId)
	if err != nil {
		return nil, nil, err
	}

	return &applicantId, &workerId, nil
}

func (h *helper) makeWorker(categoryId string, workScheduleId string) (*string, error) {
	id := uuid.NewString()
	worker := endpoints.Worker{
		Id:             id,
		Name:           "Alejandro Lopez",
		EmailAddress:   "worker@carbon-mines.pe",
		CategoryId:     categoryId,
		WorkScheduleId: workScheduleId,
	}
	if err := h.myDb.InsertWorker(&worker); err != nil {
		return nil, err
	}
	return &id, nil
}

func (h *helper) makeApplicant(categoryId string) (*string, error) {
	id := uuid.NewString()
	applicant := endpoints.Applicant{
		Id:           id,
		Name:         "Gaby Lozano",
		EmailAddress: "applicant@lozano-home.pe",
		CategoryId:   categoryId,
	}

	if err := h.myDb.InsertApplicant(&applicant); err != nil {
		return nil, err
	}

	return &id, nil
}

func databaseSetup() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("TEST_DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())
	db := endpoints.Database{
		Connection: conn,
	}
	dbHelper := helper{
		db:   conn,
		myDb: db,
	}

	applicantCatId, workerCatId, err := dbHelper.fetchCategoryIds()
	if err != nil {
		log.Fatalf("Couldn't fetch categories: %v", err)
	}
	if _, err := dbHelper.makeApplicant(*applicantCatId); err != nil {
		log.Fatalf("Couldn't populate applicant: %v", err)
	}
	workScheduleId, err := dbHelper.makeWorkSchedule()
	if err != nil {
		log.Fatalf("Couldn't populate workSchedule: %v", err)
	}
	workerId, err := dbHelper.makeWorker(*workerCatId, *workScheduleId)
	if err != nil {
		log.Fatalf("Couldn't populate workers: %v", err)
	}
	if _, err := dbHelper.makeSpecialty(*workerId); err != nil {
		log.Fatalf("Couldn't populate specialty: %v", err)
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
