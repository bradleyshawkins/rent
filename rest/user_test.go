package rest_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/bxcodec/faker/v3"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bradleyshawkins/rent"
	"github.com/bradleyshawkins/rent/rest"
	"github.com/matryer/is"
	uuid "github.com/satori/go.uuid"
)

func TestRegisterUserIntegration(t *testing.T) {
	i := is.New(t)

	user := rest.RegisterUserRequest{
		EmailAddress: faker.Email(),
		Password:     faker.Password(),
		FirstName:    faker.FirstName(),
		LastName:     faker.LastName(),
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)
	i.NoErr(err)

	req := httptest.NewRequest(http.MethodPost, "/user", &buf)
	rr := httptest.NewRecorder()

	err = router.RegisterUser(rr, req)
	i.NoErr(err)

	var rrs rest.RegisterUserResponse
	err = json.NewDecoder(rr.Body).Decode(&rrs)
	i.NoErr(err)

	i.True(rrs.UserID != uuid.Nil)
	i.True(rrs.AccountID != uuid.Nil)
}

func TestRegisterUser_EmailAddressExistsIntegration(t *testing.T) {
	i := is.New(t)

	email := faker.Email()

	user := rest.RegisterUserRequest{
		EmailAddress: email,
		Password:     faker.Password(),
		FirstName:    faker.FirstName(),
		LastName:     faker.LastName(),
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)
	i.NoErr(err)

	req := httptest.NewRequest(http.MethodPost, "/user", &buf)
	rr := httptest.NewRecorder()

	err = router.RegisterUser(rr, req)
	i.NoErr(err)

	var rrs rest.RegisterUserResponse
	err = json.NewDecoder(rr.Body).Decode(&rrs)
	i.NoErr(err)

	i.True(rrs.UserID != uuid.Nil)
	i.True(rrs.AccountID != uuid.Nil)

	dupUser := rest.RegisterUserRequest{
		EmailAddress: email,
		Password:     faker.Password(),
		FirstName:    faker.FirstName(),
		LastName:     faker.LastName(),
	}

	var dupBuf bytes.Buffer
	err = json.NewEncoder(&dupBuf).Encode(dupUser)
	i.NoErr(err)

	dupReq := httptest.NewRequest(http.MethodPost, "/user", &dupBuf)
	dupRR := httptest.NewRecorder()

	err = router.RegisterUser(dupRR, dupReq)
	if err == nil {
		t.Fatal("Should have received an error but didn't")
	}

	var rentErr *rent.Error
	if !errors.As(err, &rentErr) {
		t.Fatalf("Unexpected error type received. Error: %v", err)
	}

	t.Log(rentErr)

	if rentErr.Code() != rent.CodeDuplicate {
		t.Fatalf("unexpected code. Expected: %v, Got: %v", rent.CodeDuplicate, rentErr.Code())
	}
}

func TestRegisterUser_MissingInputIntegration(t *testing.T) {
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
			l := rest.RegisterUserRequest{
				Password:     test.password,
				FirstName:    test.firstName,
				LastName:     test.lastName,
				EmailAddress: test.emailAddress,
			}

			var buf bytes.Buffer
			err := json.NewEncoder(&buf).Encode(l)
			i.NoErr(err)

			r := httptest.NewRequest(http.MethodPost, "/user", &buf)
			rr := httptest.NewRecorder()
			err = router.RegisterUser(rr, r)
			if err == nil {
				t.Fatal("Expected an error but didn't get one")
			}

			var rentErr *rent.Error
			if !errors.As(err, &rentErr) {
				t.Fatalf("Unexpected error type received. Error: %v", err)
			}

			t.Log(rentErr)

			if rentErr.Code() != rent.CodeInvalidField {
				t.Fatalf("unexpected code. Expected: %v, Got: %v", rent.CodeInvalidField, rentErr.Code())
			}
		})
	}
}
