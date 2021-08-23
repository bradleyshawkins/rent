// +build integration

package account_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/matryer/is"

	"github.com/bradleyshawkins/rent/rest/account"
)

func TestRegisterAccount(t *testing.T) {
	i := is.New(t)
	u := os.Getenv("SERVICE_URL")
	u += "/accounts/register"
	l, err := NewRegisterAccountRequest(u, "registerAccount_register@test.com")
	i.NoErr(err)

	resp, err := http.DefaultClient.Do(l)
	i.NoErr(err)

	i.Equal(resp.StatusCode, http.StatusCreated)
}

func TestRegisterAccount_EmailAddressExists(t *testing.T) {
	i := is.New(t)

	u := os.Getenv("SERVICE_URL")
	u += "/accounts/register"
	r, err := NewRegisterAccountRequest(u, "registerAccountUsernameExists@test.com")
	i.NoErr(err)

	resp, err := http.DefaultClient.Do(r)
	i.NoErr(err)

	if resp.StatusCode != http.StatusCreated {
		b, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("expected a created status code. StatusCode: %v, Body: %v", resp.StatusCode, string(b))
	}

	r2, err := NewRegisterAccountRequest(u, "registerAccountUsernameExists@test.com")
	i.NoErr(err)

	resp2, err := http.DefaultClient.Do(r2)
	i.NoErr(err)
	if resp2.StatusCode != http.StatusConflict {
		b, _ := ioutil.ReadAll(resp2.Body)
		t.Fatalf("expected a conflict status code. StatusCode: %v, Body: %v", resp2.StatusCode, string(b))
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
			i := is.New(t)
			l := account.RegisterAccountRequest{
				Password:     test.password,
				FirstName:    test.firstName,
				LastName:     test.lastName,
				EmailAddress: test.emailAddress,
			}

			b, err := json.Marshal(l)
			i.NoErr(err)

			r, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(b))
			i.NoErr(err)

			resp, err := http.DefaultClient.Do(r)
			i.NoErr(err)

			i.Equal(resp.StatusCode, http.StatusBadRequest)
		})
	}
}

func TestRegisterPersonToAccount(t *testing.T) {
	i := is.New(t)
	u := os.Getenv("SERVICE_URL")
	registerEndpoint := u + "/accounts/register"
	l, err := NewRegisterAccountRequest(registerEndpoint, "registerPersonToAccount_success_1_register@test.com")
	i.NoErr(err)

	resp, err := http.DefaultClient.Do(l)
	i.NoErr(err)

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Unexpected status code. Expected: %v, Got: %v", http.StatusCreated, resp.StatusCode)
	}

	var acc account.RegisterAccountResponse
	err = json.NewDecoder(resp.Body).Decode(&acc)
	i.NoErr(err)

	addToAccountEndpoint := u + fmt.Sprintf("/accounts/%v/register", acc.AccountID.String())
	pReq, err := NewRegisterPersonToAccountRequest(addToAccountEndpoint, "registerPersonToAccount_success_2_register@test.com")
	i.NoErr(err)

	resp2, err := http.DefaultClient.Do(pReq)
	i.NoErr(err)

	if resp2.StatusCode != http.StatusCreated {
		b, _ := ioutil.ReadAll(resp2.Body)
		t.Fatalf("expected a created status code. StatusCode: %v, Body: %v", resp2.StatusCode, string(b))
	}
}

func TestRegisterPersonToAccount_BadInput_AccountID(t *testing.T) {
	tests := []struct {
		name       string
		accountID  string
		statusCode int
	}{
		{
			name:       "Account doesn't exist",
			accountID:  uuid.NewV4().String(),
			statusCode: http.StatusNotFound,
		},
		{
			name:       "invalid accountID",
			accountID:  "invalidAccountID",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "missing accountID",
			accountID:  "",
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := is.New(t)

			pReq := account.RegisterPersonToAccountRequest{
				Password:     "password",
				FirstName:    "firstName",
				LastName:     "lastName",
				EmailAddress: "email.address@test.com",
			}

			b, err := json.Marshal(pReq)
			i.NoErr(err)

			u := os.Getenv("SERVICE_URL")
			u += fmt.Sprintf("/accounts/%v/register", tt.accountID)
			r, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(b))
			i.NoErr(err)

			resp2, err := http.DefaultClient.Do(r)
			i.NoErr(err)

			if resp2.StatusCode != tt.statusCode {
				b, _ := ioutil.ReadAll(resp2.Body)
				t.Fatalf("expected a bad request response. StatusCode: %v, Body: %v", resp2.StatusCode, string(b))
			}
		})
	}
}

func TestRegisterPersonToAccount_BadInput_Payload(t *testing.T) {
	tests := []struct {
		name         string
		password     string
		firstName    string
		lastName     string
		emailAddress string
	}{
		{name: "Missing Password", firstName: "firstName", lastName: "lastName", emailAddress: "register_person_to_account_missing_password.address@test.com"},
		{name: "Missing FirstName", password: "password", lastName: "lastName", emailAddress: "register_person_to_account_missing_first_name.address@test.com"},
		{name: "Missing LastName", password: "password", firstName: "firstName", emailAddress: "register_person_to_account_missing_last_name@test.com"},
		{name: "Missing EmailAddress", password: "password", firstName: "firstName", lastName: "lastName"},
		{name: "Invalid EmailAddress", password: "password", firstName: "firstName", lastName: "lastName", emailAddress: "register_person_to_account_bad_email"},
	}

	for ii, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			i := is.New(t)

			u := os.Getenv("SERVICE_URL")
			registerEndpoint := u + "/accounts/register"
			l, err := NewRegisterAccountRequest(registerEndpoint, fmt.Sprintf("register_user_to_account_%v_bad_input@test.com", ii))
			i.NoErr(err)

			resp, err := http.DefaultClient.Do(l)
			i.NoErr(err)

			if resp.StatusCode != http.StatusCreated {
				b, _ := ioutil.ReadAll(resp.Body)
				t.Fatalf("failed to create new account. StatusCode: %v, Body: %v", resp.StatusCode, string(b))
			}

			var acc account.RegisterAccountResponse
			err = json.NewDecoder(resp.Body).Decode(&acc)
			i.NoErr(err)

			pReq := account.RegisterPersonToAccountRequest{
				Password:     test.password,
				FirstName:    test.firstName,
				LastName:     test.lastName,
				EmailAddress: test.emailAddress,
			}

			b, err := json.Marshal(pReq)
			i.NoErr(err)

			ur := os.Getenv("SERVICE_URL")
			ur += fmt.Sprintf("/accounts/%v/register", acc.AccountID)
			r, err := http.NewRequest(http.MethodPost, ur, bytes.NewReader(b))
			i.NoErr(err)

			resp2, err := http.DefaultClient.Do(r)
			i.NoErr(err)

			if resp2.StatusCode != http.StatusBadRequest {
				b, _ := ioutil.ReadAll(resp2.Body)
				t.Fatalf("expected bad request. StatusCode: %v, Body: %v", resp2.StatusCode, string(b))
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

func NewRegisterPersonToAccountRequest(u string, emailAddress string) (*http.Request, error) {
	b, err := json.Marshal(account.RegisterPersonToAccountRequest{
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
