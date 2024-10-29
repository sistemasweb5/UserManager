package endpoints

import (
	"api/testing/petitions"
	"net/http"
	"net/url"
)

type ApplicantTest struct {
	Address url.URL
}

func (b ApplicantTest) GetAllApplicant() (*http.Response, error) {
	return petitions.SimpleRequest(b.Address)
}

func (b ApplicantTest) GetApplicantById(id string) (*http.Response, error) {
	b.Address.Path = "client/applicant/" + id
	return petitions.SimpleRequest(b.Address)
}

func (b ApplicantTest) PostApplicant(categoryId string) (*http.Response, error) {
	body := ApplicantRequest{
		Name:         "Leonardo Lopez",
		EmailAddress: "green-bottle@applicant.com",
	}

	return petitions.BodyRequest("POST", b.Address, body)
}

func NewApplicantTest(baseUrl *string) *ApplicantTest {
	address := url.URL{
		Scheme: "http",
		Host:   *baseUrl,
		Path:   "client/applicant",
	}

	test := ApplicantTest{
		Address: address,
	}

	return &test
}
