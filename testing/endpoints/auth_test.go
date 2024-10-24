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
	path := "user/login"
	h := endpoints.NewAuthTest(&url, &path)

	log.Println("Testing SignIn with valid credentials")
	data := map[string]string{
		"email":    "test@gmail.com",
		"password": "Pass@1234123",
	}

	statusCode, _, err := h.SignIn(&data)

	if err != nil {
		t.Errorf(fmt.Sprintln(err))
	}

	if !(statusCode >= 200 && statusCode <= 299) {
		t.Errorf("Endpoint %s has failed", fmt.Sprintln(err))
	}
}

func TestAuthHandler_SignIn_Unauthorized(t *testing.T) {
	url := "localhost:5200"
	path := "user/login"
	h := endpoints.NewAuthTest(&url, &path)

	log.Println("Testing SignIn with invalid credentials")

	data := map[string]string{
		"email":    "invalid-email@gmail.com",
		"password": "Pass123",
	}

	statusCode, _, err := h.SignIn(&data)

	if err != nil {
		t.Errorf("Error during SignIn: %v", err)
		return
	}

	expectedStatusCode := http.StatusUnauthorized
	if statusCode != expectedStatusCode {
		t.Errorf("Expected status code %d, but got %d", expectedStatusCode, statusCode)
	}
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

func TestAuthHandler_SignOut_Success(t *testing.T) {
	url := "localhost:5200"
	path := "user/login"
	h := endpoints.NewAuthTest(&url, &path)

	data := map[string]string{
		"email":    "test@gmail.com",
		"password": "Pass@1234123",
	}

	_, access_token, _ := h.SignIn(&data)

	log.Println("Testing SignOut with valid token")

	path = "user/logout"
	h = endpoints.NewAuthTest(&url, &path)

	resp, err := h.SignOut(access_token)

	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Error during SignOut: %v", err)
		return
	}

	expectedStatusCode := http.StatusOK
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %d, but got %d", expectedStatusCode, resp.StatusCode)
	}
}

func TestAuthHandler_SignOut_Unauthorized(t *testing.T) {
	url := "localhost:5200"
	path := "user/logout"
	h := endpoints.NewAuthTest(&url, &path)

	log.Println("Testing SignOut with invalid token")
	invalidToken := "invalid-token"

	resp, err := h.SignOut(invalidToken)

	defer resp.Body.Close()
	if err != nil {
		t.Errorf("Error during SignOut: %v", err)
		return
	}

	expectedStatusCode := http.StatusUnauthorized
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %d, but got %d", expectedStatusCode, resp.StatusCode)
	}
}
