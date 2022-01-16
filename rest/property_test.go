package rest_test

//func TestRegisterProperty(t *testing.T) {
//	i := is.New(t)
//
//	accountID, propertyID, err := registerUserAndProperty("testRegisterProperty_test@test.com", newDefaultRegisterPropertyRequest())
//	i.NoErr(err)
//
//	i.True(accountID != (uuid.UUID{}))
//	i.True(propertyID != (uuid.UUID{}))
//}
//
//func TestRegisterProperty_BadInput(t *testing.T) {
//	street1 := "street1"
//	street2 := "street2"
//	city := "city"
//	state := "state"
//	zipcode := "zipcode"
//	accountID := uuid.NewV4().String()
//
//	tests := []struct {
//		name      string
//		street1   string
//		street2   string
//		city      string
//		state     string
//		zipcode   string
//		accountID string
//	}{
//		{
//			name:    "Missing Street1",
//			street1: "", street2: street2, city: city, state: state, zipcode: zipcode, accountID: accountID,
//		},
//		{
//			name:    "Missing City",
//			street1: street1, street2: street2, city: "", state: state, zipcode: zipcode, accountID: accountID,
//		},
//		{
//			name:    "Missing State",
//			street1: street1, street2: street2, city: city, state: "", zipcode: zipcode, accountID: accountID,
//		},
//		{
//			name:    "Missing Zipcode",
//			street1: street1, street2: street2, city: city, state: state, zipcode: "", accountID: accountID,
//		},
//	}
//
//	for idx, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			i := is.New(t)
//
//			accountID, _, err := registerUser(newDefaultRegisterUserRequest(fmt.Sprintf("testRegisterProperty_%d_test@test.com", idx)))
//			i.NoErr(err)
//
//			prop := &rest.RegisterPropertyRequest{
//				Name: tt.name,
//				Address: rest.Address{
//					Street1: tt.street1,
//					Street2: tt.street2,
//					City:    tt.city,
//					State:   tt.state,
//					Zipcode: tt.zipcode,
//				},
//			}
//
//			req, err := newRegisterPropertyRestRequest(accountID, prop)
//			i.NoErr(err)
//
//			resp, err := http.DefaultClient.Do(req)
//			i.NoErr(err)
//
//			if resp.StatusCode != http.StatusBadRequest {
//				b, _ := ioutil.ReadAll(resp.Body)
//				t.Fatalf("Unexpected status code. StatusCode: %v, Payload: %v", resp.StatusCode, string(b))
//			}
//
//			var propResp rest.Error
//			err = json.NewDecoder(resp.Body).Decode(&propResp)
//			i.NoErr(err)
//
//			i.Equal(propResp.Code, int(rent.CodeInvalidField))
//		})
//	}
//}
//
//func TestRegisterProperty_BadAccountID(t *testing.T) {
//	tests := []struct {
//		name       string
//		accountID  string
//		statusCode int
//		code       int
//	}{
//		{
//			name:       "Missing accountID",
//			accountID:  "",
//			statusCode: http.StatusBadRequest,
//			code:       int(rent.CodeInvalidField),
//		},
//		{
//			name:       "Non-UUID accountID",
//			accountID:  "1234",
//			statusCode: http.StatusBadRequest,
//			code:       int(rent.CodeInvalidField),
//		},
//		{
//			name:       "Account doesn't exist",
//			accountID:  uuid.NewV4().String(),
//			statusCode: http.StatusConflict,
//			code:       int(rent.CodeRequiredEntityNotExists),
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			i := is.New(t)
//
//			u := getServiceURL() + fmt.Sprintf("/account/%s/property", tt.accountID)
//			req, err := newRequest(http.MethodPost, u, newDefaultRegisterPropertyRequest())
//			i.NoErr(err)
//
//			resp, err := http.DefaultClient.Do(req)
//			i.NoErr(err)
//
//			err = didReceiveStatusCode(resp, tt.statusCode)
//			i.NoErr(err)
//
//			var propResp rest.Error
//			err = json.NewDecoder(resp.Body).Decode(&propResp)
//			i.NoErr(err)
//
//			i.Equal(propResp.Code, int(tt.code))
//		})
//	}
//}

//
//func TestLoadProperty(t *testing.T) {
//	i := is.New(t)
//	name := "name"
//	street1 := "street1"
//	street2 := "street2"
//	city := "city"
//	state := "state"
//	zipcode := "zipcode"
//
//	prop := &rest.RegisterPropertyRequest{
//		Name: name,
//		Address: rest.Address{
//			Street1: street1,
//			Street2: street2,
//			City:    city,
//			State:   state,
//			Zipcode: zipcode,
//		},
//	}
//
//	accountID, propertyID, err := registerUserAndProperty("testLoadProperty@test.com", prop)
//	i.NoErr(err)
//
//	u := getServiceURL() + fmt.Sprintf("/account/%s/property/%s", accountID, propertyID)
//	req, err := newRequest(http.MethodGet, u, http.NoBody)
//	i.NoErr(err)
//
//	resp, err := http.DefaultClient.Do(req)
//	i.NoErr(err)
//
//	var loadResp rest.LoadPropertyResponse
//	err = json.NewDecoder(resp.Body).Decode(&loadResp)
//	i.NoErr(err)
//
//	i.True(loadResp.ID != (uuid.UUID{}))
//	i.Equal(loadResp.Name, name)
//	i.Equal(loadResp.Address.Street1, street1)
//	i.Equal(loadResp.Address.Street2, street2)
//	i.Equal(loadResp.Address.City, city)
//	i.Equal(loadResp.Address.State, state)
//	i.Equal(loadResp.Address.Zipcode, zipcode)
//}
//
//func TestRemoveProperty(t *testing.T) {
//	i := is.New(t)
//
//	accountID, propertyID, err := registerUserAndProperty("testRemoveProperty@test.com", newDefaultRegisterPropertyRequest())
//	i.NoErr(err)
//
//	req, err := newRemovePropertyRestRequest(accountID, propertyID)
//	i.NoErr(err)
//
//	resp, err := http.DefaultClient.Do(req)
//	i.NoErr(err)
//
//	err = didReceiveStatusCode(resp, http.StatusOK)
//	i.NoErr(err)
//}
//
//func TestRemoveProperty_PropertyNotExist(t *testing.T) {
//	i := is.New(t)
//	req, err := newRemovePropertyRestRequest(uuid.NewV4(), uuid.NewV4())
//	i.NoErr(err)
//
//	resp, err := http.DefaultClient.Do(req)
//	i.NoErr(err)
//
//	err = didReceiveStatusCode(resp, http.StatusNotFound)
//	i.NoErr(err)
//}
//
//func TestRemoveProperty_BadURLParams(t *testing.T) {
//	tests := []struct {
//		name       string
//		accountID  string
//		propertyID string
//	}{
//		{
//			name:       "Non-UUID AccountID",
//			accountID:  "1234",
//			propertyID: uuid.NewV4().String(),
//		},
//		{
//			name:       "Non-UUID UserID",
//			accountID:  uuid.NewV4().String(),
//			propertyID: "1234",
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			i := is.New(t)
//			u := getServiceURL() + fmt.Sprintf("/account/%s/property/%s", tt.accountID, tt.propertyID)
//			req, err := http.NewRequest(http.MethodDelete, u, http.NoBody)
//			i.NoErr(err)
//
//			resp, err := http.DefaultClient.Do(req)
//			i.NoErr(err)
//
//			err = didReceiveStatusCode(resp, http.StatusBadRequest)
//			i.NoErr(err)
//
//			var propResp rest.Error
//			err = json.NewDecoder(resp.Body).Decode(&propResp)
//			i.NoErr(err)
//
//			i.Equal(propResp.Code, int(rent.CodeInvalidField))
//		})
//	}
//}
//
//func registerUserAndProperty(emailAddress string, prop *rest.RegisterPropertyRequest) (uuid.UUID, uuid.UUID, error) {
//	accountID, _, err := registerUser(newDefaultRegisterUserRequest(emailAddress))
//	if err != nil {
//		return uuid.UUID{}, uuid.UUID{}, err
//	}
//
//	propertyID, err := registerProperty(accountID, prop)
//	if err != nil {
//		return uuid.UUID{}, uuid.UUID{}, err
//	}
//
//	return accountID, propertyID, nil
//}
//
//func registerProperty(accountID uuid.UUID, prop *rest.RegisterPropertyRequest) (uuid.UUID, error) {
//	req, err := newRegisterPropertyRestRequest(accountID, prop)
//	if err != nil {
//		return uuid.UUID{}, err
//	}
//
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		return uuid.UUID{}, err
//	}
//
//	err = didReceiveStatusCode(resp, http.StatusCreated)
//	if err != nil {
//		return uuid.UUID{}, err
//	}
//
//	var propResp rest.RegisterPropertyResponse
//	err = json.NewDecoder(resp.Body).Decode(&propResp)
//	if err != nil {
//		return uuid.UUID{}, err
//	}
//	return propResp.PropertyID, nil
//}

//func loadProperty(accountID, propertyID uuid.UUID) (*rest.LoadPropertyResponse, error) {
//	u := getServiceURL() + fmt.Sprintf("/account/%s/property/%s", accountID, propertyID)
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
//	var loadResp rest.LoadPropertyResponse
//	err = json.NewDecoder(resp.Body).Decode(&loadResp)
//	if err != nil {
//		return nil, err
//	}
//
//	return &loadResp, nil
//}

//func newDefaultRegisterPropertyRequest() *rest.RegisterPropertyRequest {
//	return &rest.RegisterPropertyRequest{
//		Name: "test name",
//		Address: rest.Address{
//			Street1: "street1",
//			Street2: "street2",
//			City:    "city",
//			State:   "state",
//			Zipcode: "zipcode",
//		},
//	}
//}
//
//func newRegisterPropertyRestRequest(accountID uuid.UUID, prop *rest.RegisterPropertyRequest) (*http.Request, error) {
//	u := getServiceURL() + fmt.Sprintf("/account/%s/property", accountID)
//	return newRequest(http.MethodPost, u, prop)
//}
//
//func newRemovePropertyRestRequest(accountID uuid.UUID, propertyID uuid.UUID) (*http.Request, error) {
//	u := getServiceURL() + fmt.Sprintf("/account/%s/property/%s", accountID, propertyID)
//	return newRequest(http.MethodDelete, u, http.NoBody)
//}
