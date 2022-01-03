package rest

import (
	"encoding/json"
	"net/http"

	"github.com/bradleyshawkins/rent/identity"

	uuid "github.com/satori/go.uuid"

	"github.com/bradleyshawkins/rent"
)

const (
	accountID = "accountID"
	personID  = "personID"
)

type RegisterPersonRequest struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}

type RegisterPersonResponse struct {
	AccountID uuid.UUID `json:"accountID"`
	PersonID  uuid.UUID `json:"personID"`
}

func (l *Router) RegisterPerson(w http.ResponseWriter, r *http.Request) error {
	var rr RegisterPersonRequest
	err := json.NewDecoder(r.Body).Decode(&rr)
	if err != nil {
		return rent.NewError(err, rent.WithInvalidPayload(), rent.WithMessage("unable to decode request"))
	}

	person, account, err := l.registrar.Register(rr.EmailAddress, rr.FirstName, rr.LastName, rr.Password)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(RegisterPersonResponse{PersonID: person.ID.AsUUID(), AccountID: account.ID.AsUUID()})
	if err != nil {
		return rent.NewError(err, rent.WithInternal(), rent.WithMessage("unable to serialize response"))
	}

	return nil
}

type RegisterPersonToAccountRequest struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Role         string `json:"role"`
}

type RegisterPersonToAccountResponse struct {
	PersonID uuid.UUID `json:"personID"`
}

func (l *Router) RegisterPersonToAccount(w http.ResponseWriter, r *http.Request) error {
	accountID, err := getURLParamAsUUID(r, accountID)
	if err != nil {
		return err
	}

	var rr RegisterPersonToAccountRequest
	err = json.NewDecoder(r.Body).Decode(&rr)
	if err != nil {
		return rent.NewError(err, rent.WithInvalidPayload(), rent.WithMessage("unable to decode request"))
	}

	person, err := l.registrar.RegisterPersonToAccount(identity.AsAccountID(accountID), rr.Role, rr.EmailAddress, rr.FirstName, rr.LastName, rr.Password)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(RegisterPersonToAccountResponse{PersonID: person.ID.AsUUID()})
	if err != nil {
		return rent.NewError(err, rent.WithInternal(), rent.WithMessage("unable to serialize response"))
	}

	return nil
}

//
//type LoadPersonResponse struct {
//	ID           uuid.UUID `json:"id"`
//	EmailAddress string    `json:"emailAddress"`
//	FirstName    string    `json:"firstName"`
//	LastName     string    `json:"lastName"`
//	Status       int       `json:"status"`
//}
//
//func (p *Router) LoadPerson(w http.ResponseWriter, r *http.Request) error {
//	pID, err := getURLParamAsUUID(r, personID)
//	if err != nil {
//		return err
//	}
//
//	person, err := p.ps.LoadPerson(pID)
//	if err != nil {
//		return err
//	}
//
//	err = json.NewEncoder(w).Encode(LoadPersonResponse{
//		ID:           person.ID,
//		EmailAddress: person.EmailAddress,
//		FirstName:    person.FirstName,
//		LastName:     person.LastName,
//		Status:       int(person.Status),
//	})
//	if err != nil {
//		return rent.NewError(err, rent.WithInternal(), rent.WithMessage("unable to serialize get person response"))
//	}
//	return nil
//}
//
//func (p *Router) CancelPerson(w http.ResponseWriter, r *http.Request) error {
//	aID, err := getURLParamAsUUID(r, accountID)
//	if err != nil {
//		return err
//	}
//
//	pID, err := getURLParamAsUUID(r, personID)
//	if err != nil {
//		return err
//	}
//
//	err = p.ps.CancelPerson(aID, pID)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
