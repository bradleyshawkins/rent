// +build integration

package person_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/bradleyshawkins/rent/rest/person"
	"github.com/matryer/is"
	uuid "github.com/satori/go.uuid"
)

func TestLoadPerson(t *testing.T) {
	i := is.New(t)
	u := os.Getenv("SERVICE_URL")
	ea := "loadPerson_register123123123@test.com"
	fn := "test"
	ln := "user"

	registerURL := u + "/person/register"
	registerBytes, err := json.Marshal(person.RegisterPersonRequest{
		EmailAddress: ea,
		Password:     "dummyPassword",
		FirstName:    fn,
		LastName:     ln,
	})
	i.NoErr(err)

	l, err := http.NewRequest(http.MethodPost, registerURL, bytes.NewReader(registerBytes))
	i.NoErr(err)

	registerResp, err := http.DefaultClient.Do(l)
	i.NoErr(err)

	if registerResp.StatusCode != http.StatusCreated {
		b, err := ioutil.ReadAll(registerResp.Body)
		if err != nil {
			t.Fatalf("Unable to read response payload. Error: %v", err)
		}
		t.Fatalf("Unexpected status code. StatusCode: %v, Payload: %v", registerResp.StatusCode, string(b))
	}
	//
	var personResp person.RegisterPersonResponse
	err = json.NewDecoder(registerResp.Body).Decode(&personResp)
	i.NoErr(err)

	i.True(personResp.ID != (uuid.UUID{}))

	loadURL := u + "/person/" + personResp.ID.String()

	req, err := http.NewRequest(http.MethodGet, loadURL, http.NoBody)
	i.NoErr(err)

	resp, err := http.DefaultClient.Do(req)
	i.NoErr(err)

	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Unable to read response payload. Error: %v", err)
		}
		t.Fatalf("Unexpected status code. StatusCode: %v, Payload: %v", resp.StatusCode, string(b))
	}

	//var loadResp person.LoadPersonResponse
	//err = json.NewDecoder(resp.Body).Decode(&loadResp)
	//i.NoErr(err)
	//
	//i.Equal(loadResp.ID, personResp.ID)
	//i.Equal(loadResp.EmailAddress, ea)
	//i.Equal(loadResp.FirstName, fn)
	//i.Equal(loadResp.LastName, ln)
}
