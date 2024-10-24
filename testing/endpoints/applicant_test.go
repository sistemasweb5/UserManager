package endpoints_test

import (
	"api/testing/endpoints"
	"testing"
)

func TestGetAllApplicants(t *testing.T) {
	url := "localhost:8080"
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
