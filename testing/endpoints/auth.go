package endpoints

import (
	"api/testing/petitions"
	"net/http"
	"net/url"
)

type AuthTest struct {
	Address url.URL
}

type MockAuthService struct {
	Address url.URL
}

func (b MockAuthService) SignIn(data *map[string]string) (*http.Response, error) {
	return petitions.SimplePostRequest(b.Address, data)
}

func NewAuthTest(baseUrl *string) *MockAuthService {
	address := url.URL{
		Scheme: "http",
		Host:   *baseUrl,
		Path:   "user/login",
	}

	auth := MockAuthService{
		Address: address,
	}

	return &auth
}
