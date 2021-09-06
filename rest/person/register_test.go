// +build integration

package person_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/bradleyshawkins/rent/rest"

	"github.com/bradleyshawkins/rent"
	"github.com/bradleyshawkins/rent/rest/person"

	"github.com/matryer/is"
)

func TestRegisterPerson(t *testing.T) {
	i := is.New(t)
	u := os.Getenv("SERVICE_URL")
	u += "/person/register"
	l, err := NewRegisterPersonRequest(u, "registerPerson_register@test.com")
	i.NoErr(err)

	resp, err := http.DefaultClient.Do(l)
	i.NoErr(err)

	if resp.StatusCode != http.StatusCreated {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Unable to read response payload. Error: %v", err)
		}
		t.Fatalf("Unexpected status code. StatusCode: %v, Payload: %v", resp.StatusCode, string(b))
	}
}

func TestRegisterPerson_EmailAddressExists(t *testing.T) {
	i := is.New(t)

	u := os.Getenv("SERVICE_URL")
	u += "/person/register"
	r, err := NewRegisterPersonRequest(u, "registerPersonUsernameExists@test.com")
	i.NoErr(err)

	resp, err := http.DefaultClient.Do(r)
	i.NoErr(err)

	if resp.StatusCode != http.StatusCreated {
		b, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("expected a created status code. StatusCode: %v, Body: %v", resp.StatusCode, string(b))
	}

	r2, err := NewRegisterPersonRequest(u, "registerPersonUsernameExists@test.com")
	i.NoErr(err)

	resp2, err := http.DefaultClient.Do(r2)
	i.NoErr(err)
	if resp2.StatusCode != http.StatusConflict {
		b, _ := ioutil.ReadAll(resp2.Body)
		t.Fatalf("expected a conflict status code. StatusCode: %v, Body: %v", resp2.StatusCode, string(b))
	}

	var restErr rest.Error
	err = json.NewDecoder(resp2.Body).Decode(&restErr)
	i.NoErr(err)

	t.Log(restErr)

	if rent.Code(restErr.Code) != rent.CodeDuplicate {
		t.Fatalf("unexpected code. Expected: %v, Got: %v", rent.CodeDuplicate, rent.Code(restErr.Code))
	}
}

func TestRegisterPerson_BadInput(t *testing.T) {
	u := os.Getenv("SERVICE_URL")
	u += "/person/register"
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
			l := person.RegisterPersonRequest{
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

			var restErr rest.Error
			err = json.NewDecoder(resp.Body).Decode(&restErr)
			i.NoErr(err)

			t.Log(restErr)

			if rent.Code(restErr.Code) != rent.CodeInvalidField {
				t.Fatalf("unexpected code. Expected: %v, Got: %v", rent.CodeInvalidField, rent.Code(restErr.Code))
			}
		})
	}
}

func NewRegisterPersonRequest(u string, emailAddress string) (*http.Request, error) {
	b, err := json.Marshal(person.RegisterPersonRequest{
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
