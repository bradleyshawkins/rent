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

	p, err := rent.NewPerson(rr.EmailAddress, rr.Password, rr.FirstName, rr.LastName)
	if err != nil {
		return err
	}

	a := rent.NewAccount()

	err = l.ps.RegisterPerson(a, p)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(RegisterPersonResponse{PersonID: p.ID, AccountID: a.ID})
	if err != nil {
		return rent.NewError(err, rent.WithInternal(), rent.WithMessage("unable to serialize response"))
	}

	return nil
}

type LoadPersonResponse struct {
	ID           uuid.UUID `json:"id"`
	EmailAddress string    `json:"emailAddress"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Status       int       `json:"status"`
}

func (p *Router) LoadPerson(w http.ResponseWriter, r *http.Request) error {
	personID := chi.URLParam(r, personID)
	if personID == "" {
		return rent.NewError(errors.New("personID is required"), rent.WithInvalidFields(rent.InvalidField{
			Field:  "personID",
			Reason: rent.ReasonMissing,
		}), rent.WithMessage("personID is a required field"))
	}

	pID, err := uuid.FromString(personID)
	if err != nil {
		return rent.NewError(err, rent.WithInvalidFields(rent.InvalidField{
			Field:  "personID",
			Reason: rent.ReasonInvalid,
		}), rent.WithMessage("personID must be a UUID"))
	}

	person, err := p.ps.LoadPerson(pID)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(LoadPersonResponse{
		ID:           person.ID,
		EmailAddress: person.EmailAddress,
		FirstName:    person.FirstName,
		LastName:     person.LastName,
		Status:       int(person.Status),
	})
	if err != nil {
		return rent.NewError(err, rent.WithInternal(), rent.WithMessage("unable to serialize get person response"))
	}
	return nil
}
