package account_test

//func TestPersonEndpoints(t *testing.T) {
//	i := is.New(t)
//	u := os.Getenv("SERVICE_URL")
//	u += "/person"
//	l, err := NewRegisterAccountRequest(u, "registerPerson_all", "registerPerson_all@test.com")
//	i.NoErr(err)
//
//	resp, err := http.DefaultClient.Do(l)
//	i.NoErr(err)
//
//	i.True(resp.StatusCode == http.StatusCreated)
//
//	var registerResponse account.RegisterAccountResponse
//	err = json.NewDecoder(resp.Body).Decode(&registerResponse)
//	i.NoErr(err)
//
//}
