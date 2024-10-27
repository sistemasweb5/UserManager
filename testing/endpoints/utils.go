package endpoints

import (
	"io"
	"net/http"
	"testing"
)

func Helper(t *testing.T, resp *http.Response){
	statusCode := resp.StatusCode
	if !(statusCode >= 200 && statusCode <= 299) {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}
		
		t.Errorf("Response status code outside expected range\nStatus code: %d\nResponse body: %v", statusCode, string(body))
	}
}

