package endpoints_test

import (
	"api/testing/endpoints"
	"os"
	"testing"
)

func TestGetAllApplicants(t *testing.T) {
	url := os.Getenv("USER_MANAGER_HOSTNAME")
	endpoint := endpoints.NewApplicantTest(&url)
	resp, err := endpoint.GetAll()
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Could not reach endpoint %s", endpoint.Address.String())
	}

	statusCode := resp.StatusCode
	if !(statusCode >= 200 && statusCode <= 299) {
		t.Errorf("Endpoint %s has failed", endpoint.Address.String())
	}
}
