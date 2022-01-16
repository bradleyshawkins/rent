package rest_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bradleyshawkins/rent/identity"
	"github.com/bradleyshawkins/rent/location"
	"github.com/bradleyshawkins/rent/postgres"
	"github.com/bradleyshawkins/rent/rest"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

var (
	router *rest.Router
)

func TestMain(m *testing.M) {
	// Create Database
	log.Println("Beginning integration tests")
	db, err := postgres.NewDatabase("postgresql://postgres:password@localhost:5432/rent?sslmode=disable", "../postgres/schema")
	if err != nil {
		log.Println("Unable to create database connection. Error:", err)
		os.Exit(999)
	}
	// Create registrar
	r := identity.NewRegistrar(db)
	// Create location creator
	l := location.NewCreator(db)
	// Create Router
	router = rest.NewRouter(r, l)

	code := m.Run()

	log.Println("Completed integration tests")
	os.Exit(code)
}

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
