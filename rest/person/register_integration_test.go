// +build integration

package person_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/bradleyshawkins/rent/rest/person"
)

func TestRegisterLandlord(t *testing.T) {
	u := os.Getenv("SERVICE_URL")
	u += "/landlord"
	l, err := NewRegisterLandlordRequest(u, "registerLandlord_register", "registerLandlord_register@test.com")
	if err != nil {
		t.Fatalf("Unable to create landlord. Error: %v", err)
	}

	resp, err := http.DefaultClient.Do(l)
	if err != nil {
		t.Fatalf("Unable to make http request. Url: %v Error: %v", u, err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Unexpected status code. Expected: %v, Got: %v", http.StatusCreated, resp.StatusCode)
	}
}

func TestRegisterLandlord_UsernameExists(t *testing.T) {
	u := os.Getenv("SERVICE_URL")
	u += "/landlord"
	r, err := NewRegisterLandlordRequest(u, "registerLandlordUsernameExists", "registerLandlordUsernameExists@test.com")
	if err != nil {
		t.Fatalf("Unable to create landlord request")
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Fatalf("Unable to make http request. Url: %v Error: %v", u, err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Unexpected status code. Expected: %v, Got: %v", http.StatusCreated, resp.StatusCode)
	}

	r2, err := NewRegisterLandlordRequest(u, "registerLandlordUsernameExists", "registerLandlordUsernameExists@test.com")
	if err != nil {
		t.Fatalf("Unable to create landlord request. Error: %v", err)
	}

	resp2, err := http.DefaultClient.Do(r2)
	if err != nil {
		t.Fatalf("Unable to make http request. Url: %v Error: %v", u, err)
	}

	if resp2.StatusCode != http.StatusConflict {
		t.Errorf("Unexpected status code. Expected: %v, Got: %v", http.StatusConflict, resp.StatusCode)
	}
}

func TestRegisterLandlord_BadInput(t *testing.T) {
	u := os.Getenv("SERVICE_URL")
	u += "/landlord"
	tests := []struct {
		name         string
		username     string
		password     string
		firstName    string
		lastName     string
		emailAddress string
	}{
		{name: "Missing Username", password: "password", firstName: "firstName", lastName: "lastName", emailAddress: "test.address@test.com"},
		{name: "Missing Password", username: "username", firstName: "firstName", lastName: "lastName", emailAddress: "test.address@test.com"},
		{name: "Missing FirstName", username: "username", password: "password", lastName: "lastName", emailAddress: "test.address@test.com"},
		{name: "Missing LastName", username: "username", password: "password", firstName: "firstName", emailAddress: "test.address@test.com"},
		{name: "Missing EmailAddress", username: "username", password: "password", firstName: "firstName", lastName: "lastName"},
		{name: "Invalid EmailAddress", username: "username", password: "password", firstName: "firstName", lastName: "lastName", emailAddress: "test"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := person.RegisterRequest{
				Username:     test.username,
				Password:     test.password,
				FirstName:    test.firstName,
				LastName:     test.lastName,
				EmailAddress: test.emailAddress,
			}

			b, err := json.Marshal(l)
			if err != nil {
				t.Fatalf("Unable to marshal landlord register request. Error: %v", err)
			}

			r, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(b))
			if err != nil {
				t.Fatalf("Unable to create http request. Url: %v Error: %v", u, err)
			}

			resp, err := http.DefaultClient.Do(r)
			if err != nil {
				t.Fatalf("Unable to make http request. Url: %v Error: %v", u, err)
			}

			if resp.StatusCode != http.StatusBadRequest {
				t.Errorf("Unexpected status code. Expected: %v, Got: %v", http.StatusBadRequest, resp.StatusCode)
			}
		})
	}

}

func NewRegisterLandlordRequest(u string, username string, emailAddress string) (*http.Request, error) {
	b, err := json.Marshal(person.RegisterRequest{
		Username:     username,
		Password:     "password",
		FirstName:    "FirstName",
		LastName:     "LastName",
		EmailAddress: emailAddress,
	})
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	return r, nil
}
