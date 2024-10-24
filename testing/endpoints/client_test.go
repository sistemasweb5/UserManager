package endpoints_test

import (
	"api/testing/endpoints"
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func setup() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("TEST_DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	db := endpoints.Database{
		Connection: conn,
	}

	categoryId := uuid.NewString()
	client := endpoints.Client{
		Id:           uuid.NewString(),
		Name:         "Slim shady",
		EmailAddress: "fake_email@mail.com",
		CategoryId:   categoryId,
	}
	category := endpoints.Category{
		Id:  categoryId,
		Rol: "Plumber",
	}

	if err := db.InsertClient(&client); err != nil {
		log.Fatalf("Could not populate category table: %v", err)
	}
	if err := db.InsertCategory(&category); err != nil {
		log.Fatalf("Could not populate client table: %v", err)
	}
}

func TestMain(m *testing.M) {
	log.SetFlags(log.Lshortfile)
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestGetAll(t *testing.T) {
	url := "localhost:5200"
	client := endpoints.NewClientTest(&url)
	resp, err := client.GetAll()
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Could not reach endpoint %s", client.Address.String())
	}

	statusCode := resp.StatusCode
	if !(statusCode >= 200 && statusCode <= 299) {
		t.Errorf("Endpoint %s has failed", client.Address.String())
	}
}

func TestGetById(t *testing.T) {
	url := "localhost:5200"
	clientEndpoint := endpoints.NewClientTest(&url)

	respClients, err := clientEndpoint.GetAll()
	defer respClients.Body.Close()
	body, err := io.ReadAll(respClients.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	var clients []endpoints.Client
	if err := json.Unmarshal(body, &clients); err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(clients) < 1 {
		t.Error("There was no clients")
	}
	client := clients[0]

	resp, err := clientEndpoint.GetById(client.Id)
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Could not reach endpoint %s", clientEndpoint.Address.String())
	}
	statusCode := resp.StatusCode
	if !(statusCode >= 200 && statusCode <= 299) {
		t.Errorf("Endpoint %s has failed", clientEndpoint.Address.String())
	}
}

func TestFetchInvalidId(t *testing.T) {
	url := "localhost:5200"
	clientEndpoint := endpoints.NewClientTest(&url)
	invalidID := "00000000-0000-0000-0000-000000000000"
	resp, err := clientEndpoint.GetById(invalidID)
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Could not reach endpoint %s", clientEndpoint.Address.String())
	}
	statusCode := resp.StatusCode
	if statusCode >= 200 && statusCode <= 299 {
		t.Errorf("Endpoint %s has failed", clientEndpoint.Address.String())
	}
}
