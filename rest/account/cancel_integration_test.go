package account_test

import (
	"testing"
)

func TestCancel(t *testing.T) {
	//u := os.Getenv("SERVICE_URL")
	//u += "/person"
	//
	//l, err := NewRegisterAccountRequest(u, "registerPerson-cancel", "registerPerson_cancel@test.com")
	//if err != nil {
	//	t.Fatalf("Unable to create person. Error: %v", err)
	//}
	//
	//resp, err := http.DefaultClient.Do(l)
	//if err != nil {
	//	t.Fatalf("Unable to make http request. Url: %v Error: %v", u, err)
	//}
	//
	//if resp.StatusCode != http.StatusCreated {
	//	t.Fatalf("Unexpected status code. Expected: %v, Got: %v", http.StatusCreated, resp.StatusCode)
	//}
	//
	//var registerResponse account.RegisterAccountResponse
	//err = json.NewDecoder(resp.Body).Decode(&registerResponse)
	//if err != nil {
	//	t.Fatalf("unable to decode register person response. Error: %v", err)
	//}
	//
	//cr, err := http.NewRequest(http.MethodDelete, u+"/"+registerResponse.PersonID.String(), http.NoBody)
	//if err != nil {
	//	t.Fatalf("unable to create delete person request. Error: %v", err)
	//}
	//
	//cancelResp, err := http.DefaultClient.Do(cr)
	//if err != nil {
	//	t.Fatalf("unable to make request to delete person. Error: %v", err)
	//}
	//
	//if cancelResp.StatusCode != http.StatusOK {
	//	t.Errorf("unexpected status code. Expected: %v, Got: %v", http.StatusOK, cancelResp.StatusCode)
	//}
}
