package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"

	uuid "github.com/satori/go.uuid"

	"github.com/bradleyshawkins/rent"
)

const (
	propertyID = "propertyID"
)

type RegisterPropertyRequest struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

type RegisterPropertyResponse struct {
	PropertyID uuid.UUID `json:"propertyID"`
}

type Address struct {
	Street1 string `json:"street1"`
	Street2 string `json:"street2"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode string `json:"zipcode"`
}

func (l *Router) RegisterProperty(w http.ResponseWriter, r *http.Request) error {
	accID := chi.URLParam(r, accountID)
	if accID == "" {
		return rent.NewError(errors.New("accountID is required"), rent.WithInvalidFields(rent.InvalidField{
			Field:  "accountID",
			Reason: rent.ReasonMissing,
		}), rent.WithMessage("accountID is a required field"))
	}

	aID, err := uuid.FromString(accID)
	if err != nil {
		return rent.NewError(err, rent.WithInvalidFields(rent.InvalidField{
			Field:  "accountID",
			Reason: rent.ReasonInvalid,
		}), rent.WithMessage("accountID must be a UUID"))
	}

	var rr RegisterPropertyRequest
	err = json.NewDecoder(r.Body).Decode(&rr)
	if err != nil {
		return rent.NewError(err, rent.WithInvalidPayload(), rent.WithMessage("unable to decode register payload request"))
	}

	addr, err := rent.NewAddress(rr.Address.Street1, rr.Address.Street2, rr.Address.City, rr.Address.State, rr.Address.Zipcode)
	if err != nil {
		return err
	}

	prop, err := rent.NewProperty(rr.Name, addr)
	if err != nil {
		return err
	}

	err = l.propStore.RegisterProperty(aID, prop)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(RegisterPropertyResponse{PropertyID: prop.ID})
	if err != nil {
		return rent.NewError(err, rent.WithInternal(), rent.WithMessage("unable to serialize response"))
	}
	return nil
}

type LoadPropertyResponse struct {
	ID      uuid.UUID
	Name    string
	Address Address
}

func (l *Router) LoadProperty(w http.ResponseWriter, r *http.Request) error {
	aID, err := getURLParamAsUUID(r, accountID)
	if err != nil {
		return err
	}

	pID, err := getURLParamAsUUID(r, propertyID)
	if err != nil {
		return err
	}

	prop, err := l.propStore.LoadProperty(aID, pID)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(LoadPropertyResponse{
		ID:   prop.ID,
		Name: prop.Name,
		Address: Address{
			Street1: prop.Address.Street1,
			Street2: prop.Address.Street2,
			City:    prop.Address.City,
			State:   prop.Address.State,
			Zipcode: prop.Address.Zipcode,
		},
	})

	return nil
}

func (l *Router) RemoveProperty(w http.ResponseWriter, r *http.Request) error {
	aID, err := getURLParamAsUUID(r, accountID)
	if err != nil {
		return err
	}

	pID, err := getURLParamAsUUID(r, propertyID)
	if err != nil {
		return err
	}

	err = l.propStore.RemoveProperty(aID, pID)
	if err != nil {
		return err
	}

	return nil
}
