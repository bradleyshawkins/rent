// +build integration

package http_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	uuid "github.com/satori/go.uuid"

	h "github.com/bradleyshawkins/rent/http"
)

func TestRegisterPerson(t *testing.T) {
	person := &h.Person{
		FirstName:    "Bradley",
		LastName:     "Hawkins",
		EmailAddress: "bradleyshawkins@gmail.com",
	}

	personID, err := insertPerson(person)
	if err != nil {
		t.Fatalf("unable to insert person. Error: %v", err)
	}

	person.ID = personID

	insertedPerson, err := getPerson(personID)
	if err != nil {
		t.Fatalf("unable to get person. Error: %v", err)
	}

	if *insertedPerson != *person {
		t.Errorf("Unexpected Person. Expected: %v, Got: %v", person, insertedPerson)
	}
}

func TestGetUserNotExist(t *testing.T) {
	u := fmt.Sprintf("http://localhost:8080/person/%v", uuid.NewV4())
	req, err := http.NewRequest(http.MethodGet, u, http.NoBody)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("unable to make request to rent service. Error: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Unexpected status code. Expected: %v, Got: %v", http.StatusNotFound, resp.StatusCode)
	}
}

func TestCreateGetUpdateDeletePerson(t *testing.T) {
	initialPerson := &h.Person{
		FirstName:    "Brad",
		LastName:     "Hawkins",
		EmailAddress: "bradleyshawkins@gmail.com",
	}

	id, err := insertPerson(initialPerson)
	if err != nil {
		t.Fatalf("unable to insert person. Error: %v", err)
	}

	initialPerson.ID = id

	p, err := getPerson(id)
	if err != nil {
		t.Fatalf("unable to get person. Error: %v", err)
	}

	if *initialPerson != *p {
		t.Errorf("Retrieved person was not the same as the inserted person. Inserted: %v, Got: %v", initialPerson, p)
	}

	u := &h.Person{
		ID:           id,
		FirstName:    "Bradley",
		LastName:     "Hawkins",
		EmailAddress: "brad.hawkins@gmail.com",
	}

	err = updatePerson(u)
	if err != nil {
		t.Fatalf("Unable to update person. Error: %v", err)
	}

	p, err = getPerson(id)
	if err != nil {
		t.Fatalf("Unable to get updated person. Error: %v", err)
	}

	if *p != *u {
		t.Errorf("Retrieved person was not the same as the updated person. Updated: %v, Got: %v", u, p)
	}

	err = deletePerson(id)
	if err != nil {
		t.Errorf("Unable to delete person. Error: %v", err)
	}

	wasDeleted, err := verifyDoesNotExist(id)
	if err != nil {
		t.Fatalf("unable to verify that person does not exist. Error: %v", err)
	}

	if !wasDeleted {
		t.Errorf("Person was not deleted")
	}

}

func insertPerson(p *h.Person) (string, error) {
	resp, err := doRequest(http.MethodPost, "http://localhost:8080/register", p)
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
	resp, err := doRequest(http.MethodGet, u, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("unable to make request to get person. Error: %v", err)
	}

	defer resp.Body.Close()

	var p h.Person
	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response person. Error:%v", err)
	}

	return &p, nil
}

func updatePerson(p *h.Person) error {
	u := fmt.Sprintf("http://localhost:8080/person/%v", p.ID)
	resp, err := doRequest(http.MethodPut, u, p)
	if err != nil {
		return fmt.Errorf("unable to update person. Error: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code. Status Code: %v", resp.StatusCode)
	}
	return nil
}

func deletePerson(id string) error {
	u := fmt.Sprintf("http://localhost:8080/person/%v", id)
	resp, err := doRequest(http.MethodDelete, u, http.NoBody)
	if err != nil {
		return fmt.Errorf("unable to update person. Error: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code. Status Code: %v", resp.StatusCode)
	}
	return nil
}

func verifyDoesNotExist(id string) (bool, error) {
	u := fmt.Sprintf("http://localhost:8080/person/%v", id)
	req, err := http.NewRequest(http.MethodGet, u, http.NoBody)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("unable to make request. Error: %v", err)
	}

	defer resp.Body.Close()

	return resp.StatusCode == http.StatusNotFound, nil
}

func doRequest(method string, u string, body interface{}) (*http.Response, error) {

	for i := 0; i < 3; i++ {
		var requestBody io.Reader = http.NoBody
		if body != nil {
			b, _ := json.Marshal(body)
			requestBody = bytes.NewReader(b)
		}

		req, err := http.NewRequest(method, u, requestBody)
		if err != nil {
			return nil, fmt.Errorf("unable to create request Error: %v", err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("unable to make request. Error: %v", err)
		}

		if resp.StatusCode > http.StatusMultipleChoices {
			resp.Body.Close()
			return nil, fmt.Errorf("got a bad status code. Status Code: %v", resp.StatusCode)
		}

		return resp, nil
	}

	return nil, errors.New("unable to make request after 3 times")
}
