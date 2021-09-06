package person

import (
	"encoding/json"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/bradleyshawkins/rent"
)

type RegisterPersonRequest struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}

type RegisterPersonResponse struct {
	ID uuid.UUID `json:"id"`
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

	err = l.ps.RegisterPerson(p)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(RegisterPersonResponse{ID: p.ID})
	if err != nil {
		return rent.NewError(err, rent.WithInternal(), rent.WithMessage("unable to serialize response"))
	}

	return nil
}
