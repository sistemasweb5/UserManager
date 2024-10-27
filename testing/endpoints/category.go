package endpoints

import (
	"api/testing/petitions"
	"net/http"
	"net/url"
)

type CategoryTest struct {
	Address url.URL
}

func (b CategoryTest) GetAllCategory() (*http.Response, error) {
	return petitions.SimpleRequest(b.Address)
}

func NewCategoryTest(baseUrl *string) *CategoryTest {
	address := url.URL{
		Scheme: "http",
		Host:   *baseUrl,
		Path:   "client/category",
	}

	client := CategoryTest{
		Address: address,
	}

	return &client
}
