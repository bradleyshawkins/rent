package rest_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getServiceURL() string {
	u := os.Getenv("SERVICE_URL")
	if u == "" {
		u = "http://127.0.0.1:8080"
	}
	return u
}

func newRequest(method string, u string, payload interface{}) (*http.Request, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func didReceiveStatusCode(resp *http.Response, expected int) error {
	if resp.StatusCode != expected {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("unable to read response payload. Error: %w", err)
		}
		return fmt.Errorf("unexpected status code. Expected: %d, StatusCode: %d, Payload: %s", expected, resp.StatusCode, string(b))
	}
	return nil
}
