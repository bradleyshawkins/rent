package person_test

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/bradleyshawkins/rent/rest/person"
	"github.com/matryer/is"
)

func TestLandlordEndpoints(t *testing.T) {
	i := is.New(t)
	u := os.Getenv("SERVICE_URL")
	u += "/landlord"
	l, err := NewRegisterLandlordRequest(u, "registerLandlord_all", "registerLandlord_all@test.com")
	i.NoErr(err)

	resp, err := http.DefaultClient.Do(l)
	i.NoErr(err)

	i.True(resp.StatusCode == http.StatusCreated)

	var registerResponse person.RegisterResponse
	err = json.NewDecoder(resp.Body).Decode(&registerResponse)
	i.NoErr(err)

}
