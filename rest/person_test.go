//go:build integration
// +build integration

package rest_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bradleyshawkins/rent"
	"github.com/bradleyshawkins/rent/rest"
	"github.com/matryer/is"
	uuid "github.com/satori/go.uuid"
)

func TestRegisterPerson(t *testing.T) {
	i := is.New(t)

	accountID, personID, err := registerPerson(newDefaultRegisterPersonRequest("registerPerson_register@test.com"))
	i.NoErr(err)

	i.True(accountID != (uuid.UUID{}))
	i.True(personID != (uuid.UUID{}))
}

func TestRegisterPerson_EmailAddressExists(t *testing.T) {
	i := is.New(t)

	u := getServiceURL()
	u += "/person/register"

	_, _, err := registerPerson(newDefaultRegisterPersonRequest("registerPersonUsernameExists@test.com"))
	i.NoErr(err)

	r2, err := NewRegisterPersonRequest(u, "registerPersonUsernameExists@test.com")
	i.NoErr(err)

	resp2, err := http.DefaultClient.Do(r2)
	i.NoErr(err)

	err = didReceiveStatusCode(resp2, http.StatusConflict)
	i.NoErr(err)

	var restErr rest.Error
	err = json.NewDecoder(resp2.Body).Decode(&restErr)
	i.NoErr(err)

	t.Log(restErr)

	if rent.Code(restErr.Code) != rent.CodeDuplicate {
		t.Fatalf("unexpected code. Expected: %v, Got: %v", rent.CodeDuplicate, rent.Code(restErr.Code))
	}
}

func TestRegisterPerson_BadInput(t *testing.T) {
	u := getServiceURL()
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
			l := rest.RegisterPersonRequest{
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
	b, err := json.Marshal(rest.RegisterPersonRequest{
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

//
//func TestLoadPerson(t *testing.T) {
//	i := is.New(t)
//	u := getServiceURL()
//	ea := "loadPerson_register@test.com"
//	fn := "test"
//	ln := "user"
//
//	registerURL := u + "/person/register"
//	registerBytes, err := json.Marshal(rest.RegisterPersonRequest{
//		EmailAddress: ea,
//		Password:     "dummyPassword",
//		FirstName:    fn,
//		LastName:     ln,
//	})
//	i.NoErr(err)
//
//	l, err := http.NewRequest(http.MethodPost, registerURL, bytes.NewReader(registerBytes))
//	i.NoErr(err)
//
//	registerResp, err := http.DefaultClient.Do(l)
//	i.NoErr(err)
//
//	err = didReceiveStatusCode(registerResp, http.StatusCreated)
//	i.NoErr(err)
//
//	var personResp rest.RegisterPersonResponse
//	err = json.NewDecoder(registerResp.Body).Decode(&personResp)
//	i.NoErr(err)
//
//	i.True(personResp.PersonID != (uuid.UUID{}))
//
//	loadPersonResp, err := loadPerson(personResp.PersonID)
//	i.NoErr(err)
//
//	i.True(loadPersonResp != nil)
//
//	i.Equal(loadPersonResp.ID, personResp.PersonID)
//	i.Equal(loadPersonResp.EmailAddress, ea)
//	i.Equal(loadPersonResp.FirstName, fn)
//	i.Equal(loadPersonResp.LastName, ln)
//}
//
//func TestCancelPerson(t *testing.T) {
//	i := is.New(t)
//	accountID, personID, err := registerPerson(newDefaultRegisterPersonRequest("registerPerson_cancel@test.com"))
//	i.NoErr(err)
//
//	u := getServiceURL() + fmt.Sprintf("/account/%s/person/%s", accountID.String(), personID.String())
//	r, err := http.NewRequest(http.MethodDelete, u, http.NoBody)
//	i.NoErr(err)
//
//	resp, err := http.DefaultClient.Do(r)
//	i.NoErr(err)
//	defer resp.Body.Close()
//
//	err = didReceiveStatusCode(resp, http.StatusOK)
//	i.NoErr(err)
//
//	getURL := getServiceURL() + "/person/" + personID.String()
//	req, err := http.NewRequest(http.MethodGet, getURL, http.NoBody)
//	i.NoErr(err)
//
//	loadResp, err := http.DefaultClient.Do(req)
//	i.NoErr(err)
//	defer loadResp.Body.Close()
//
//	i.NoErr(didReceiveStatusCode(loadResp, http.StatusNotFound))
//}
//
//func TestCancelPerson_BadInput(t *testing.T) {
//	tests := []struct {
//		name       string
//		accountID  string
//		personID   string
//		statusCode int
//		code       int
//	}{
//		{
//			name:       "Invalid AccountID",
//			accountID:  "1234",
//			personID:   uuid.NewV4().String(),
//			statusCode: http.StatusBadRequest,
//			code:       int(rent.CodeInvalidField),
//		},
//		{
//			name:       "Invalid PersonID",
//			accountID:  uuid.NewV4().String(),
//			personID:   "1234",
//			statusCode: http.StatusBadRequest,
//			code:       int(rent.CodeInvalidField),
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			i := is.New(t)
//
//			u := getServiceURL() + fmt.Sprintf("/account/%s/person/%s", tt.accountID, tt.personID)
//			r, err := http.NewRequest(http.MethodDelete, u, http.NoBody)
//			i.NoErr(err)
//
//			resp, err := http.DefaultClient.Do(r)
//			i.NoErr(err)
//			defer resp.Body.Close()
//
//			err = didReceiveStatusCode(resp, tt.statusCode)
//			i.NoErr(err)
//
//			var restErr rest.Error
//			err = json.NewDecoder(resp.Body).Decode(&restErr)
//			i.NoErr(err)
//
//			i.True(restErr.Code == tt.code)
//		})
//	}
//}
//
//func TestCancelPerson_PersonNotExist(t *testing.T) {
//	i := is.New(t)
//	u := getServiceURL() + fmt.Sprintf("/account/%s/person/%s", uuid.NewV4().String(), uuid.NewV4().String())
//	r, err := http.NewRequest(http.MethodDelete, u, http.NoBody)
//	i.NoErr(err)
//
//	resp, err := http.DefaultClient.Do(r)
//	i.NoErr(err)
//	defer resp.Body.Close()
//
//	err = didReceiveStatusCode(resp, http.StatusNotFound)
//	i.NoErr(err)
//
//	var restErr rest.Error
//	err = json.NewDecoder(resp.Body).Decode(&restErr)
//	i.NoErr(err)
//
//	i.True(restErr.Code == int(rent.CodeNotExists))
//}

func registerPerson(p *rest.RegisterPersonRequest) (uuid.UUID, uuid.UUID, error) {
	req, err := newRegisterPersonRestRequest(p)
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, err
	}

	registerPersonResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, err
	}

	err = didReceiveStatusCode(registerPersonResp, http.StatusCreated)
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, err
	}

	var personResp rest.RegisterPersonResponse
	err = json.NewDecoder(registerPersonResp.Body).Decode(&personResp)
	if err != nil {
		return uuid.UUID{}, uuid.UUID{}, err
	}
	return personResp.AccountID, personResp.PersonID, nil
}

//func loadPerson(pID uuid.UUID) (*rest.LoadPersonResponse, error) {
//	u := getServiceURL() + "/person/" + pID.String()
//
//	req, err := http.NewRequest(http.MethodGet, u, http.NoBody)
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		return nil, err
//	}
//
//	err = didReceiveStatusCode(resp, http.StatusOK)
//	if err != nil {
//		return nil, err
//	}
//
//	var loadResp rest.LoadPersonResponse
//	err = json.NewDecoder(resp.Body).Decode(&loadResp)
//	if err != nil {
//		return nil, err
//	}
//
//	return &loadResp, nil
//}

func newRegisterPersonRestRequest(r *rest.RegisterPersonRequest) (*http.Request, error) {
	u := getServiceURL() + "/person/register"
	return newRequest(http.MethodPost, u, r)
}

func newDefaultRegisterPersonRequest(ea string) *rest.RegisterPersonRequest {
	return &rest.RegisterPersonRequest{
		EmailAddress: ea,
		Password:     "password",
		FirstName:    "firstName",
		LastName:     "lastName",
	}
}
