package person

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/bradleyshawkins/rent"
	"github.com/bradleyshawkins/rent/rest"
)

type RegisterRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
}

type RegisterResponse struct {
	PersonID uuid.UUID `json:"personID"`
}

func (l *Router) Register(w http.ResponseWriter, r *http.Request) {
	if err := l.register(w, r); err != nil {
		err.WriteError(w)
		return
	}
}

func (l *Router) register(w http.ResponseWriter, r *http.Request) *rest.Error {
	var rr RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&rr)
	if err != nil {
		log.Println(fmt.Errorf("unable to decode request. Error: %v", err))
		return rest.NewError(http.StatusBadRequest, "invalid request")
	}

	p, err := rent.NewPerson(rr.Username, rr.Password, rr.FirstName, rr.LastName, rr.EmailAddress, "")
	if err != nil {
		log.Println(fmt.Errorf("invalid person creation. Error: %v", err))
		return rest.NewError(http.StatusBadRequest, err.Error())
	}

	if err := l.personService.RegisterPerson(p); err != nil {
		log.Println(fmt.Errorf("unable to create person. personID: %s Error: %v", p.ID, err))
		if rest.IsDuplicate(err) {
			return rest.NewError(http.StatusConflict, "username already exists")
		}
		return rest.NewError(http.StatusInternalServerError, "unable to update user")
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(RegisterResponse{PersonID: p.ID}); err != nil {
		log.Print(fmt.Errorf("unable to marshal person creation. Error: %v", err))
		return rest.NewError(http.StatusInternalServerError, "unable to marshal landlord response")
	}

	return nil
}
