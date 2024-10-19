package petitions

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

func marshalBody(body *any) *[]byte {
	jsonBody, err := json.Marshal(&body)
	if err != nil {
		log.Fatalf("Could not marshal book body: %v", err)
	}

	return &jsonBody
}

func BodyRequest(httpMethod string, postURL url.URL, body any) (*http.Response, error) {
	log.Printf("Processing %s request: %s", httpMethod, postURL.String())
	jsonBody := marshalBody(&body)
	bodyReader := bytes.NewReader(*jsonBody)

	request, err := http.NewRequest(httpMethod, postURL.String(), bodyReader)
	if err != nil {
		log.Printf("Could not create %s request: %v", httpMethod, err)
		return nil, err
	}

	if httpMethod != http.MethodPut || httpMethod != http.MethodPost {
		request.Header.Set("Content-Type", "application/json")
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Failed to execute %s request: %v", httpMethod, err)
		return nil, err
	}

	return resp, nil
}

func SimpleRequest(postURL url.URL) (*http.Response, error) {
	log.Printf("Processing GET request: %s", postURL.String())
	resp, err := http.Get(postURL.String())
	if err != nil {
		log.Printf("An error occured during processing: %v", err)
		return nil, err
	}

	return resp, nil
}

