package landlord_test

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/bradleyshawkins/rent/http/landlord"
)

func TestLandlordEndpoints(t *testing.T) {
	u := os.Getenv("SERVICE_URL")
	u += "/landlord"
	l, err := NewRegisterLandlordRequest(u, "registerLandlord_all", "registerLandlord_all@test.com")
	if err != nil {
		t.Fatalf("Unable to create landlord. Error: %v", err)
	}

	resp, err := http.DefaultClient.Do(l)
	if err != nil {
		t.Fatalf("Unable to make http request. Url: %v Error: %v", u, err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Unexpected status code. Expected: %v, Got: %v", http.StatusCreated, resp.StatusCode)
	}

	var registerResponse landlord.RegisterResponse
	err = json.NewDecoder(resp.Body).Decode(&registerResponse)
	if err != nil {
		t.Fatalf("unable to decode register landlord response. Error: %v", err)
	}

}
