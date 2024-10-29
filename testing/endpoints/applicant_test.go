package endpoints_test

import (
	"api/testing/endpoints"
	"encoding/json"
	"io"
	"log"
	"os"
	"testing"
)

func TestGetAllApplicant(t *testing.T) {
	url := os.Getenv("USER_MANAGER_HOSTNAME")
	endpoint := endpoints.NewApplicantTest(&url)
	resp, err := endpoint.GetAllApplicant()
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Could not reach endpoint %s", endpoint.Address.String())
	}

	endpoints.Helper(t, resp)
}

func TestGetAllApplicantIsNotEmpty(t *testing.T) {
	url := os.Getenv("USER_MANAGER_HOSTNAME")
	endpoint := endpoints.NewApplicantTest(&url)

	respClients, err := endpoint.GetAllApplicant()
	defer respClients.Body.Close()
	body, err := io.ReadAll(respClients.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	var workers []endpoints.ApplicantResponse
	if err := json.Unmarshal(body, &workers); err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(workers) == 0 {
		t.Errorf("Expected at least one item, but endpoint returned 0\n Response: %v", workers)
	}
}

func TestGetApplicantById(t *testing.T) {
	url := os.Getenv("USER_MANAGER_HOSTNAME")
	endpoint := endpoints.NewApplicantTest(&url)

	respClients, err := endpoint.GetAllApplicant()
	defer respClients.Body.Close()
	body, err := io.ReadAll(respClients.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	var array []endpoints.ApplicantResponse
	if err := json.Unmarshal(body, &array); err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(array) < 1 {
		t.Error("Response body is empty")
	}
	item := array[0]

	resp, err := endpoint.GetApplicantById(item.Id)
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Could not reach endpoint %s", endpoint.Address.String())
	}
	statusCode := resp.StatusCode
	if !(statusCode >= 200 && statusCode <= 299) {
		t.Errorf("Response status code outside expected range\nStatus code: %d", statusCode)
		log.Print(array)
	}
}

func TestPostApplicant(t *testing.T){
	url := os.Getenv("USER_MANAGER_HOSTNAME")
	endpoint := endpoints.NewApplicantTest(&url)
	resp, err := endpoint.PostApplicant("3c89fd37-0976-4afc-846d-d87cd4b589b0")
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Could not reach endpoint %s", endpoint.Address.String())
	}

	endpoints.Helper(t, resp)
}
