package endpoints_test

import (
	"api/testing/endpoints"
	"testing"
)

func TestGetAll(t *testing.T) {
	url := "localhost:5200"
	client := endpoints.NewClientTest(&url)
	statusCode, err  := client.GetAll()
	if err != nil {
		t.Errorf("Could not reach endpoint %s", client.Address.String())
	}

	if !(statusCode >= 200 && statusCode <= 299) {
		t.Errorf("Endpoint %s has failed", client.Address.String())
	}
}

func TestGetById(t *testing.T) {
	url := "localhost:5200"
	id := "aaaaaaaa-1111-1111-1111-111111111111"
	client := endpoints.NewClientTest(&url)
	statusCode, err  := client.GetById(id)
	if err != nil {
		t.Errorf("Could not reach endpoint %s", client.Address.String())
	}

	if !(statusCode >= 200 && statusCode <= 299) {
		t.Errorf("Endpoint %s has failed", client.Address.String())
	}
}
