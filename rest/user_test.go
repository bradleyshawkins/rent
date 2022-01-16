package rest_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bxcodec/faker/v3"

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

	req, err := http.NewRequest(http.MethodPost, serverAddr+"/users", &buf)
	i.NoErr(err)

	resp, err := httpClient.Do(req)
	i.NoErr(err)

	i.Equal(resp.StatusCode, http.StatusCreated)

	var rrs rest.RegisterUserResponse
	err = json.NewDecoder(resp.Body).Decode(&rrs)
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

	req, err := http.NewRequest(http.MethodPost, serverAddr+"/users", &buf)
	i.NoErr(err)

	resp, err := httpClient.Do(req)
	i.NoErr(err)

	var rrs rest.RegisterUserResponse
	err = json.NewDecoder(resp.Body).Decode(&rrs)
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

	dupReq, err := http.NewRequest(http.MethodPost, serverAddr+"/users", &dupBuf)
	i.NoErr(err)

	dupResp, err := httpClient.Do(dupReq)
	i.NoErr(err)

	i.Equal(dupResp.StatusCode, http.StatusConflict)

	var restErr *rest.Error
	err = json.NewDecoder(dupResp.Body).Decode(&restErr)
	i.NoErr(err)
	t.Log(restErr)

	i.Equal(restErr.Code, int(rent.CodeDuplicate))
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

			r, err := http.NewRequest(http.MethodPost, serverAddr+"/users", &buf)
			i.NoErr(err)

			resp, err := httpClient.Do(r)
			i.NoErr(err)

			i.Equal(resp.StatusCode, http.StatusBadRequest)

			var restErr *rest.Error
			err = json.NewDecoder(resp.Body).Decode(&restErr)
			i.NoErr(err)
			t.Log(restErr)

			i.Equal(restErr.Code, int(rent.CodeInvalidField))
		})
	}
}

func TestLoadPersonIntegration(t *testing.T) {
	i := is.New(t)

	// Create user
	user := rest.RegisterUserRequest{
		EmailAddress: faker.Email(),
		Password:     faker.Password(),
		FirstName:    faker.FirstName(),
		LastName:     faker.LastName(),
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)
	i.NoErr(err)

	req, err := http.NewRequest(http.MethodPost, serverAddr+"/users", &buf)
	i.NoErr(err)

	resp, err := httpClient.Do(req)
	i.NoErr(err)

	var rrs rest.RegisterUserResponse
	err = json.NewDecoder(resp.Body).Decode(&rrs)
	i.NoErr(err)

	i.True(rrs.UserID != uuid.Nil)
	i.True(rrs.AccountID != uuid.Nil)

	// Load User
	addr := serverAddr + "/users/" + rrs.UserID.String()
	loadReq, err := http.NewRequest(http.MethodGet, addr, http.NoBody)
	i.NoErr(err)

	loadResp, err := httpClient.Do(loadReq)
	i.NoErr(err)

	var lu rest.LoadUserResponse
	err = json.NewDecoder(loadResp.Body).Decode(&lu)
	i.NoErr(err)

	i.Equal(lu.ID, rrs.UserID)
	i.Equal(lu.FirstName, user.FirstName)
	i.Equal(lu.LastName, user.LastName)
	i.Equal(lu.EmailAddress, user.EmailAddress)
}

func TestLoadUser_UserNotExistIntegration(t *testing.T) {
	i := is.New(t)

	addr := serverAddr + "/users/" + uuid.NewV4().String()
	loadReq, err := http.NewRequest(http.MethodGet, addr, http.NoBody)
	i.NoErr(err)

	resp, err := httpClient.Do(loadReq)
	i.NoErr(err)

	i.Equal(resp.StatusCode, http.StatusNotFound)

	var restErr *rest.Error
	err = json.NewDecoder(resp.Body).Decode(&restErr)
	i.NoErr(err)
	t.Log(restErr)

	i.Equal(restErr.Code, int(rent.CodeNotExists))
}
