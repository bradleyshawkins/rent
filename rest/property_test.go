// +build integration

package rest_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/bradleyshawkins/rent"

	uuid "github.com/satori/go.uuid"

	"github.com/bradleyshawkins/rent/rest"
	"github.com/matryer/is"
)

func TestRegisterProperty(t *testing.T) {
	i := is.New(t)
	u := getServiceURL()

	accountID, _, err := registerPerson("testRegisterProperty_test@test.com")
	i.NoErr(err)

	rPropU := u + fmt.Sprintf("/account/%s/property", accountID)
	rpr, err := newRegisterPropertyRequest(rPropU)
	i.NoErr(err)

	resp, err := http.DefaultClient.Do(rpr)
	i.NoErr(err)

	if resp.StatusCode != http.StatusCreated {
		b, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("Unexpected status code. StatusCode: %v, Payload: %v", resp.StatusCode, string(b))
	}

	var propResp rest.RegisterPropertyResponse
	err = json.NewDecoder(resp.Body).Decode(&propResp)
	i.NoErr(err)

	i.True(propResp.PropertyID != (uuid.UUID{}))
}

func TestRegisterProperty_BadInput(t *testing.T) {
	street1 := "street1"
	street2 := "street2"
	city := "city"
	state := "state"
	zipcode := "zipcode"
	accountID := uuid.NewV4().String()

	tests := []struct {
		name      string
		street1   string
		street2   string
		city      string
		state     string
		zipcode   string
		accountID string
	}{
		{
			name:    "Missing Street1",
			street1: "", street2: street2, city: city, state: state, zipcode: zipcode, accountID: accountID,
		},
		{
			name:    "Missing City",
			street1: street1, street2: street2, city: "", state: state, zipcode: zipcode, accountID: accountID,
		},
		{
			name:    "Missing State",
			street1: street1, street2: street2, city: city, state: "", zipcode: zipcode, accountID: accountID,
		},
		{
			name:    "Missing Zipcode",
			street1: street1, street2: street2, city: city, state: state, zipcode: "", accountID: accountID,
		},
	}

	for idx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := is.New(t)
			u := getServiceURL()

			accountID, _, err := registerPerson(fmt.Sprintf("testRegisterProperty_%d_test@test.com", idx))
			i.NoErr(err)

			rPropU := u + fmt.Sprintf("/account/%s/property", accountID)
			prop := rest.RegisterPropertyRequest{
				Name: tt.name,
				Address: rest.Address{
					Street1: tt.street1,
					Street2: tt.street2,
					City:    tt.city,
					State:   tt.state,
					Zipcode: tt.zipcode,
				},
			}

			b, err := json.Marshal(prop)
			i.NoErr(err)

			rpr, err := http.NewRequest(http.MethodPost, rPropU, bytes.NewBuffer(b))
			i.NoErr(err)

			resp, err := http.DefaultClient.Do(rpr)
			i.NoErr(err)

			if resp.StatusCode != http.StatusBadRequest {
				b, _ := ioutil.ReadAll(resp.Body)
				t.Fatalf("Unexpected status code. StatusCode: %v, Payload: %v", resp.StatusCode, string(b))
			}

			var propResp rest.Error
			err = json.NewDecoder(resp.Body).Decode(&propResp)
			i.NoErr(err)

			i.Equal(propResp.Code, int(rent.CodeInvalidField))
		})
	}
}

func TestRegisterProperty_BadAccountID(t *testing.T) {
	tests := []struct {
		name       string
		accountID  string
		statusCode int
		code       int
	}{
		{
			name:       "Missing accountID",
			accountID:  "",
			statusCode: http.StatusBadRequest,
			code:       int(rent.CodeInvalidField),
		},
		{
			name:       "Non-UUID accountID",
			accountID:  "1234",
			statusCode: http.StatusBadRequest,
			code:       int(rent.CodeInvalidField),
		},
		{
			name:       "Account doesn't exist",
			accountID:  uuid.NewV4().String(),
			statusCode: http.StatusConflict,
			code:       int(rent.CodeRequiredEntityNotExists),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := is.New(t)
			u := getServiceURL()

			rPropU := u + fmt.Sprintf("/account/%s/property", tt.accountID)
			rpr, err := newRegisterPropertyRequest(rPropU)
			i.NoErr(err)

			resp, err := http.DefaultClient.Do(rpr)
			i.NoErr(err)

			if resp.StatusCode != tt.statusCode {
				b, _ := ioutil.ReadAll(resp.Body)
				t.Fatalf("Unexpected status code. Expected: %v, Got: %v, Payload: %v", tt.statusCode, resp.StatusCode, string(b))
			}

			var propResp rest.Error
			err = json.NewDecoder(resp.Body).Decode(&propResp)
			i.NoErr(err)

			i.Equal(propResp.Code, int(tt.code))
		})
	}
}

func newRegisterPropertyRequest(u string) (*http.Request, error) {
	b, err := json.Marshal(rest.RegisterPropertyRequest{
		Name: "Test Register Property",
		Address: rest.Address{
			Street1: "street1",
			Street2: "street2",
			City:    "city",
			State:   "state",
			Zipcode: "zipcode",
		},
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return req, nil
}
