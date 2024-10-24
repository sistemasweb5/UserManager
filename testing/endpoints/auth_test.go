package endpoints_test

import (
	"api/testing/endpoints"
	"fmt"
	"net/http"
	"testing"

	"log"
)

func TestAuthHandler_SignIn_Success(t *testing.T) {
	url := "localhost:5200"
	h := endpoints.NewAuthTest(&url)

	log.Println("Testing SignIn with valid credentials")
	data := map[string]string{
		"email":    "safaportafolios@gmail.com",
		"password": "Pass@1234123",
	}

	resp, err := h.SignIn(&data)

	defer resp.Body.Close()
	if err != nil {
		t.Errorf(fmt.Sprintln(err))
	}

	statusCode := resp.StatusCode
	if !(statusCode >= 200 && statusCode <= 299) {
		t.Errorf("Endpoint %s has failed", fmt.Sprintln(err))
	}
}

func TestAuthHandler_SignIn_Unauthorized(t *testing.T) {
	url := "localhost:5200"
	h := endpoints.NewAuthTest(&url)

	log.Println("Testing SignIn with invalid credentials")

	data := map[string]string{
		"email":    "invalid-email@gmail.com",
		"password": "Pass123",
	}

	resp, err := h.SignIn(&data)

	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Error during SignIn: %v", err)
		return
	}

	expectedStatusCode := http.StatusUnauthorized
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %d, but got %d", expectedStatusCode, resp.StatusCode)
	}
}
