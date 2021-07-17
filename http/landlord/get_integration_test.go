package landlord_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/matryer/is"

	"github.com/bradleyshawkins/rent/http/landlord"
)

func TestGetLandlord(t *testing.T) {
	i := is.New(t)
	u := os.Getenv("SERVICE_URL")
	u += "/landlord"

	reg := landlord.RegisterRequest{
		Username:     "registerRequest",
		Password:     "registerPassword",
		FirstName:    "firstName",
		LastName:     "lastName",
		EmailAddress: "first.last@test.com",
	}
	b, err := json.Marshal(reg)
	i.NoErr(err)

	l, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(b))
	i.NoErr(err)

	resp, err := http.DefaultClient.Do(l)
	i.NoErr(err)

	i.True(resp.StatusCode == http.StatusCreated)

	var registerResponse landlord.RegisterResponse
	err = json.NewDecoder(resp.Body).Decode(&registerResponse)
	i.NoErr(err)

	cr, err := http.NewRequest(http.MethodGet, u+"/"+registerResponse.LandlordID.String(), http.NoBody)
	i.NoErr(err)

	getResp, err := http.DefaultClient.Do(cr)
	i.NoErr(err)

	i.True(getResp.StatusCode == http.StatusOK)

	var gl landlord.GetLandlordResponse
	err = json.NewDecoder(getResp.Body).Decode(&gl)
	i.NoErr(err)

	i.True(registerResponse.LandlordID == gl.LandlordID)
	i.True(gl.LastName == reg.LastName)
	i.True(gl.FirstName == reg.FirstName)
}
