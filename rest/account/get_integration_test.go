package account_test

import (
	"testing"
)

func TestGetPerson(t *testing.T) {
	//i := is.New(t)
	//u := os.Getenv("SERVICE_URL")
	//u += "/person"
	//
	//reg := account.RegisterAccountRequest{
	//	Username:     "registerRequest",
	//	Password:     "registerPassword",
	//	FirstName:    "firstName",
	//	LastName:     "lastName",
	//	EmailAddress: "first.last@test.com",
	//}
	//b, err := json.Marshal(reg)
	//i.NoErr(err)
	//
	//l, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(b))
	//i.NoErr(err)
	//
	//resp, err := http.DefaultClient.Do(l)
	//i.NoErr(err)
	//
	//i.True(resp.StatusCode == http.StatusCreated)
	//
	//var registerResponse account.RegisterAccountResponse
	//err = json.NewDecoder(resp.Body).Decode(&registerResponse)
	//i.NoErr(err)
	//
	//cr, err := http.NewRequest(http.MethodGet, u+"/"+registerResponse.PersonID.String(), http.NoBody)
	//i.NoErr(err)
	//
	//getResp, err := http.DefaultClient.Do(cr)
	//i.NoErr(err)
	//
	//i.True(getResp.StatusCode == http.StatusOK)
	//
	//var gl account.GetPersonResponse
	//err = json.NewDecoder(getResp.Body).Decode(&gl)
	//i.NoErr(err)
	//
	//i.True(registerResponse.PersonID == gl.PersonID)
	//i.True(gl.LastName == reg.LastName)
	//i.True(gl.FirstName == reg.FirstName)
}
