package endpoints

import (
	"api/testing/petitions"
	"net/http"
	"net/url"
)


type ApplicantTest struct {
	Address url.URL
}

func (b ApplicantTest) GetAll() (*http.Response, error) {
	return petitions.SimpleRequest(b.Address)
}

func NewApplicantTest(baseUrl *string) *ClientTest {
	address := url.URL{
		Scheme: "http",
		Host:   *baseUrl,
		Path:   "client/applicant",
	}

	client := ClientTest{
		Address: address,
	}

	return &client
}
