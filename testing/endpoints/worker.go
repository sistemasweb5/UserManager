package endpoints

import (
	"api/testing/petitions"
	"net/http"
	"net/url"
)

type WorkerTest struct {
	Address url.URL
}

func (b WorkerTest) GetWorkerById(id string) (*http.Response, error) {
	b.Address.Path = "client/worker/" + id
	return petitions.SimpleRequest(b.Address)
}

func (b WorkerTest) GetAllWorkers() (*http.Response, error) {
	return petitions.SimpleRequest(b.Address)
}

func (b WorkerTest) PostWorker(categoryId string) (*http.Response, error) {
	body := WorkerRequest{
		Name:           "Posted applicant",
		EmailAddress:   "worker@mine-os.com",
		WorkScheduleId: "d688c312-c735-43d2-9d38-312168856c21",
	}

	return petitions.BodyRequest("POST", b.Address, body)
}

func NewWorkerTest(baseUrl *string) *WorkerTest {
	address := url.URL{
		Scheme: "http",
		Host:   *baseUrl,
		Path:   "client/worker",
	}

	client := WorkerTest{
		Address: address,
	}

	return &client
}
