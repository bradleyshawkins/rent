package person

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bradleyshawkins/rent"

	h "github.com/bradleyshawkins/rent/http"
)

type RegisterRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"` // TODO: Use better password handling
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
	PhoneNumber  string `json:"phoneNumber"`
}

type RegisterResponse struct {
	ID string `json:"id"`
}

func (p *Router) RegisterPerson(w http.ResponseWriter, r *http.Request) {
	if err := p.registerPerson(w, r); err != nil {
		err.WriteError(w)
	}
}

func (p *Router) registerPerson(w http.ResponseWriter, r *http.Request) *h.Error {
	var rr RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&rr)
	if err != nil {
		log.Println(fmt.Errorf("unable to decode user. Error: %v", err))
		return h.NewError(http.StatusBadRequest, "unable to unmarshal request")
	}

	person, err := rent.NewPerson(rr.Username, rr.Password, rr.FirstName, rr.LastName, rr.EmailAddress, rr.PhoneNumber)
	if err != nil {
		log.Println(fmt.Errorf("error creating person. Error: %v", err))
		return h.NewError(http.StatusBadRequest, "missing field")
	}

	id, err := p.personService.CreatePerson(person)
	if err != nil {
		log.Println(fmt.Errorf("unable to update user. userID: %s Error: %v", person.ID, err))
		if h.IsDuplicate(err) {
			return h.NewError(http.StatusConflict, "username already exists")
		}
		return h.NewError(http.StatusInternalServerError, "unable to update user")
	}

	response := RegisterResponse{
		ID: id.String(),
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(fmt.Errorf("unable to marshal response. Error: %v", err))
		return h.NewError(http.StatusInternalServerError, "unable to marshal response")
	}

	return nil
}
