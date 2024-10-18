package petitions

import (
	"bytes"
	"encoding/json"
	"io"
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

func BodyRequest(httpMethod string, postURL url.URL, body any) (int, error) {
	log.Printf("Processing %s request: %s", httpMethod, postURL.String())
	jsonBody := marshalBody(&body)
	bodyReader := bytes.NewReader(*jsonBody)

	request, err := http.NewRequest(httpMethod, postURL.String(), bodyReader)
	if err != nil {
		log.Printf("Could not create %s request: %v", httpMethod, err)
		return -1, err
	}

	if httpMethod != http.MethodPut || httpMethod != http.MethodPost {
		request.Header.Set("Content-Type", "application/json")
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Failed to execute %s request: %v", httpMethod, err)
		return -1, err
	}
	defer resp.Body.Close()

	if statusCode, err := handleRequest(httpMethod, *resp); err != nil {
		return -1, err
	} else {
		return statusCode, nil
	}
}

func SimpleRequest(postURL url.URL) (int, error) {
	log.Printf("Processing GET request: %s", postURL.String())
	resp, err := http.Get(postURL.String())
	if err != nil {
		log.Printf("An error occured during processing: %v", err)
		return -1, err
	}
	defer resp.Body.Close()

	statusCode, err := handleRequest("GET", *resp)
	if err != nil {
		log.Printf("Failed to execute %s request: %v", "GET", err)
		return -1, err
	}

	return statusCode, err
}

func handleRequest(httpMethod string, resp http.Response) (int, error) {
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}
	defer resp.Body.Close()
	log.Print(resp.Status)
	log.Println("Response:\n", string(respBody))

	return resp.StatusCode, nil
}
