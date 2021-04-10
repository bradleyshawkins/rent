// +build integration

package http_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	h "github.com/bradleyshawkins/rent/http"
)

func TestRegisterPerson(t *testing.T) {
	person := h.Person{
		FirstName:     "Bradley",
		MiddleInitial: "S",
		LastName:      "Hawkins",
	}

	personID, err := insertPerson(person)
	if err != nil {
		t.Fatalf("unable to insert person. Error: %v", err)
	}

	insertedPerson, err := getPerson(personID)
	if err != nil {
		t.Fatalf("unable to get person. Error: %v", err)
	}

	if person.FirstName != insertedPerson.FirstName {
		t.Errorf("Unexpected First Name. Expected: %v, Got: %v", person.FirstName, insertedPerson.FirstName)
	}

	if person.MiddleInitial != insertedPerson.MiddleInitial {
		t.Errorf("Unexpected Middle Initial. Expected: %v, Got: %v", person.MiddleInitial, insertedPerson.MiddleInitial)
	}

	if person.LastName != insertedPerson.LastName {
		t.Errorf("Unexpected Last Name. Expected: %v, Got: %v", person.LastName, insertedPerson.LastName)
	}
}

func insertPerson(p h.Person) (string, error) {
	b, _ := json.Marshal(p)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/register", bytes.NewReader(b))
	if err != nil {
		return "", fmt.Errorf("Unable to create request to create person Error: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("unable to make request to create person. Error: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("unexpected status code. Expected: %v, Got: %v", http.StatusCreated, resp.StatusCode)
	}

	var response h.RegisterResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", fmt.Errorf("unable to decode response. Error: %v", err)
	}
	return response.ID, nil
}

func getPerson(id string) (*h.Person, error) {
	u := fmt.Sprintf("http://localhost:8080/person/%v", id)
	req, err := http.NewRequest(http.MethodGet, u, http.NoBody)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Unable to make request to get person. Error: %v", err)
	}

	defer resp.Body.Close()

	var p h.Person
	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response person. Error:%v", err)
	}

	return &p, nil
}
