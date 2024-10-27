package endpoints

import (
	"api/testing/petitions"
	"net/http"
	"net/url"
)

type UserTest struct {
	Address url.URL
}

func (b UserTest) GetAllApplicant() (*http.Response, error) {
	b.Address.Path = "client/applicant"
	return petitions.SimpleRequest(b.Address)
}

func (b UserTest) GetAllWorker() (*http.Response, error) {
	b.Address.Path = "client/worker"
	return petitions.SimpleRequest(b.Address)
}

func (b UserTest) GetById(id string) (*http.Response, error) {
	b.Address.Path = "client/" + id
	return petitions.SimpleRequest(b.Address)
}

func NewClientTest(baseUrl *string) *UserTest {
	address := url.URL{
		Scheme: "http",
		Host:   *baseUrl,
		Path:   "user/",
	}

	client := UserTest{
		Address: address,
	}

	return &client
}
