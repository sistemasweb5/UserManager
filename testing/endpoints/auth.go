package endpoints

import (
	"api/testing/petitions"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type AuthTest struct {
	Address url.URL
}

type MockAuthService struct {
	Address url.URL
}

func (b MockAuthService) SignIn(data *map[string]string) (int, string, error) {
	resp, err := petitions.SimplePostRequest(b.Address, data, "")
	if err != nil {
		return 0, "", fmt.Errorf("error during SignIn: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, "", fmt.Errorf("error reading response body: %v", err)
	}

	var authResp struct {
		AccessToken string `json:"access_token"`
	}
	err = json.Unmarshal(respBody, &authResp)
	if err != nil {
		return resp.StatusCode, "", fmt.Errorf("error unmarshalling response body: %v", err)
	}

	return resp.StatusCode, authResp.AccessToken, nil
}

func (b MockAuthService) SignOut(accessToken string) (*http.Response, error) {
	return petitions.SimplePostRequest(b.Address, nil, accessToken)
}

func NewAuthTest(baseUrl *string, path *string) *MockAuthService {
	address := url.URL{
		Scheme: "http",
		Host:   *baseUrl,
		Path:   *path,
	}

	auth := MockAuthService{
		Address: address,
	}

	return &auth
}
