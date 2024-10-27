package endpoints_test

import (
	"api/testing/endpoints"
	"encoding/json"
	"io"
	"log"
	"os"
	"testing"
)

func TestGetAllCategory(t *testing.T) {
	url := os.Getenv("USER_MANAGER_HOSTNAME")
	endpoint := endpoints.NewCategoryTest(&url)
	resp, err := endpoint.GetAllCategory()
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Could not reach endpoint %s", endpoint.Address.String())
	}

	statusCode := resp.StatusCode
	if !(statusCode >= 200 && statusCode <= 299) {
		t.Errorf("Response status code outside expected range\nStatus code: %d", statusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Could not retrieve body: %v", err)
	}

	var categories []endpoints.Category
	if err := json.Unmarshal(body, &categories); err != nil {
		t.Errorf("Error: %v", err)
	}
	if size := len(categories); size != 2 {
		t.Fatalf("Expected 2 categories, found %d", size)
	}

	var workerCategory endpoints.Category
	var applicantCategory endpoints.Category

	workerCategory = categories[0]
	if workerCategory.Rol != "worker" {
		workerCategory = categories[1]
		applicantCategory = categories[0]
	}

	if workerCategory.Rol != "worker" {
		t.Fatalf("Expected category worker, got %s", workerCategory.Rol)
	}

	if applicantCategory.Rol != "applicant" {
		t.Fatalf("Expected category applicant, got %s", applicantCategory.Rol)
	}
}
