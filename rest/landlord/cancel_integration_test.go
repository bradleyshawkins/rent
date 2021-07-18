package landlord_test

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/bradleyshawkins/rent/rest/landlord"
)

func TestCancel(t *testing.T) {
	u := os.Getenv("SERVICE_URL")
	u += "/landlord"

	l, err := NewRegisterLandlordRequest(u, "registerLandlord-cancel", "registerLandlord_cancel@test.com")
	if err != nil {
		t.Fatalf("Unable to create landlord. Error: %v", err)
	}

	resp, err := http.DefaultClient.Do(l)
	if err != nil {
		t.Fatalf("Unable to make http request. Url: %v Error: %v", u, err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Unexpected status code. Expected: %v, Got: %v", http.StatusCreated, resp.StatusCode)
	}

	var registerResponse landlord.RegisterResponse
	err = json.NewDecoder(resp.Body).Decode(&registerResponse)
	if err != nil {
		t.Fatalf("unable to decode register landlord response. Error: %v", err)
	}

	cr, err := http.NewRequest(http.MethodDelete, u+"/"+registerResponse.LandlordID.String(), http.NoBody)
	if err != nil {
		t.Fatalf("unable to create delete landlord request. Error: %v", err)
	}

	cancelResp, err := http.DefaultClient.Do(cr)
	if err != nil {
		t.Fatalf("unable to make request to delete landlord. Error: %v", err)
	}

	if cancelResp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code. Expected: %v, Got: %v", http.StatusOK, cancelResp.StatusCode)
	}
}
