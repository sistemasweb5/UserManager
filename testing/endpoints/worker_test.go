package endpoints_test

import (
	"api/testing/endpoints"
	"encoding/json"
	"io"
	"os"
	"testing"
)

func TestGetAllWorker(t *testing.T) {
	url := os.Getenv("USER_MANAGER_HOSTNAME")
	endpoint := endpoints.NewWorkerTest(&url)
	resp, err := endpoint.GetAllWorkers()
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Could not reach endpoint %s", endpoint.Address.String())
	}
	endpoints.Helper(t, resp)
}

func TestGetAllWorkerIsNotEmpty(t *testing.T) {
	url := os.Getenv("USER_MANAGER_HOSTNAME")
	clientEndpoint := endpoints.NewClientTest(&url)

	respClients, err := clientEndpoint.GetAllWorker()
	defer respClients.Body.Close()
	body, err := io.ReadAll(respClients.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	var workers []endpoints.WorkerResponse
	if err := json.Unmarshal(body, &workers); err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(workers) == 0 {
		t.Errorf("Expected at least one item, but endpoint returned 0\n Response: %v", workers)
	}
}

func TestGetWorkerById(t *testing.T) {
	url := os.Getenv("USER_MANAGER_HOSTNAME")
	clientEndpoint := endpoints.NewWorkerTest(&url)

	respClients, err := clientEndpoint.GetAllWorkers()
	defer respClients.Body.Close()
	body, err := io.ReadAll(respClients.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	var workers []endpoints.WorkerResponse
	if err := json.Unmarshal(body, &workers); err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(workers) < 1 {
		t.Error("Response body is empty")
	}
	worker := workers[0]

	resp, err := clientEndpoint.GetWorkerById(worker.Id)
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Could not reach endpoint %s", clientEndpoint.Address.String())
	}
	statusCode := resp.StatusCode
	if !(statusCode >= 200 && statusCode <= 299) {
		t.Errorf("Response status code outside expected range\nStatus code: %d", statusCode)
	}
}

func TestPostWorker(t *testing.T){
	url := os.Getenv("USER_MANAGER_HOSTNAME")
	endpoint := endpoints.NewWorkerTest(&url)
	resp, err := endpoint.PostWorker("3c89fd37-0976-4afc-846d-d87cd4b589b0")
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Could not reach endpoint %s", endpoint.Address.String())
	}

	endpoints.Helper(t, resp)
}
