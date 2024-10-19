package endpoints

import (
	"api/testing/petitions"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	Id           string
	Name         string
	EmailAddress string
	CategoryId   string
}

type Category struct {
	Id  string
	Rol string
}

type ClientTest struct {
	Address url.URL
}

func (b ClientTest) GetAll() (*http.Response, error) {
	return petitions.SimpleRequest(b.Address)
}

func (b ClientTest) GetById(id string) (*http.Response, error) {
	b.Address.Path = "client/" + id
	return petitions.SimpleRequest(b.Address)
}

func NewClientTest(baseUrl *string) *ClientTest {
	address := url.URL{
		Scheme: "http",
		Host:   *baseUrl,
		Path:   "client",
	}

	client := ClientTest{
		Address: address,
	}

	return &client
}

func InitClient(baseUrl *string) {
	client := NewClientTest(baseUrl)
	log.Printf("Testing client endpoints")
	client.GetAll()
	client.GetById("aaaaaaaa-1111-1111-1111-111111111111")
}
