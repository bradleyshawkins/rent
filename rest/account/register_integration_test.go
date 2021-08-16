// +build integration

package account_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/bradleyshawkins/rent/rest/account"
)

func TestRegisterAccount(t *testing.T) {
	u := os.Getenv("SERVICE_URL")
	u += "/accounts/register"
	l, err := NewRegisterAccountRequest(u, "registerAccount_register@test.com")
	if err != nil {
		t.Fatalf("Unable to create person. Error: %v", err)
	}

	resp, err := http.DefaultClient.Do(l)
	if err != nil {
		t.Fatalf("Unable to make http request. Url: %v Error: %v", u, err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Unexpected status code. Expected: %v, Got: %v", http.StatusCreated, resp.StatusCode)
	}
}

func TestRegisterAccount_EmailAddressExists(t *testing.T) {
	u := os.Getenv("SERVICE_URL")
	u += "/accounts/register"
	r, err := NewRegisterAccountRequest(u, "registerAccountUsernameExists@test.com")
	if err != nil {
		t.Fatalf("Unable to create person request")
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Fatalf("Unable to make http request. Url: %v Error: %v", u, err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Unexpected status code. Expected: %v, Got: %v", http.StatusCreated, resp.StatusCode)
	}

	r2, err := NewRegisterAccountRequest(u, "registerAccountUsernameExists@test.com")
	if err != nil {
		t.Fatalf("Unable to create person request. Error: %v", err)
	}

	resp2, err := http.DefaultClient.Do(r2)
	if err != nil {
		t.Fatalf("Unable to make http request. Url: %v Error: %v", u, err)
	}

	if resp2.StatusCode != http.StatusConflict {
		t.Errorf("Unexpected status code. Expected: %v, Got: %v", http.StatusConflict, resp.StatusCode)
	}
}

func TestRegisterAccount_BadInput(t *testing.T) {
	u := os.Getenv("SERVICE_URL")
	u += "/accounts/register"
	tests := []struct {
		name         string
		password     string
		firstName    string
		lastName     string
		emailAddress string
	}{
		{name: "Missing Password", firstName: "firstName", lastName: "lastName", emailAddress: "test.address@test.com"},
		{name: "Missing FirstName", password: "password", lastName: "lastName", emailAddress: "test.address@test.com"},
		{name: "Missing LastName", password: "password", firstName: "firstName", emailAddress: "test.address@test.com"},
		{name: "Missing EmailAddress", password: "password", firstName: "firstName", lastName: "lastName"},
		{name: "Invalid EmailAddress", password: "password", firstName: "firstName", lastName: "lastName", emailAddress: "test"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := account.RegisterAccountRequest{
				Password:     test.password,
				FirstName:    test.firstName,
				LastName:     test.lastName,
				EmailAddress: test.emailAddress,
			}

			b, err := json.Marshal(l)
			if err != nil {
				t.Fatalf("Unable to marshal person register request. Error: %v", err)
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

func NewRegisterAccountRequest(u string, emailAddress string) (*http.Request, error) {
	b, err := json.Marshal(account.RegisterAccountRequest{
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
