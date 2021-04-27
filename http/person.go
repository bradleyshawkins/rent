package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/bradleyshawkins/rent"
	"github.com/go-chi/chi"
)

const personID = "personID"

type Person struct {
	ID           string `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
}

type RegisterResponse struct {
	ID string `json:"id"`
}

func (r *Router) Register(w http.ResponseWriter, req *http.Request) {
	var person Person
	err := json.NewDecoder(req.Body).Decode(&person)
	if err != nil {
		log.Println(fmt.Errorf("unable to decode Person. Error: %v", err))
		NewError(http.StatusBadRequest, "unable to unmarshal request").WriteError(w)
		return
	}

	p := &rent.Person{
		FirstName:    person.FirstName,
		LastName:     person.LastName,
		EmailAddress: person.EmailAddress,
	}

	id, err := r.personService.Register(p)
	if err != nil {
		log.Println(fmt.Errorf("unable to update Person. PersonID: %s Error: %v", p.ID, err))
		if isDuplicate(err) {
			NewError(http.StatusConflict, "email address already exists")
			return
		}
		NewError(http.StatusInternalServerError, "unable to update Person").WriteError(w)
		return
	}

	response := RegisterResponse{
		ID: id.String(),
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(fmt.Errorf("unable to marshal response. Error: %v", err))
		NewError(http.StatusInternalServerError, "unable to marshal response")
		return
	}
}

func (r *Router) GetPerson(w http.ResponseWriter, req *http.Request) {
	log.Println("Starting get person request")
	pID := chi.URLParam(req, personID)
	if pID == "" {
		log.Println("personID was not provided")
		NewError(http.StatusBadRequest, "personID is required").WriteError(w)
		return
	}

	id, err := uuid.FromString(pID)
	if err != nil {
		log.Println("id was not a uuid. ID:", pID)
		NewError(http.StatusBadRequest, "id was not a uuid").WriteError(w)
		return
	}

	p, err := r.personService.GetPerson(id)
	if err != nil {
		log.Println(fmt.Errorf("unable to get Person by id: %s Error: %v", pID, err))
		NewError(http.StatusInternalServerError, "unable to get Person by ID").WriteError(w)
		return
	}

	if p == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	respTenant := Person{
		ID:           p.ID.String(),
		FirstName:    p.FirstName,
		LastName:     p.LastName,
		EmailAddress: p.EmailAddress,
	}

	fmt.Printf("%+v\n", respTenant)

	err = json.NewEncoder(w).Encode(respTenant)
	if err != nil {
		log.Println(fmt.Errorf("unable to marshal tenant tenantID: %s, Error: %v", pID, err))
		NewError(http.StatusInternalServerError, "unable to marshal tenant").WriteError(w)
		return
	}
}

func (r *Router) UpdatePerson(w http.ResponseWriter, req *http.Request) {
	pID := chi.URLParam(req, personID)
	if pID == "" {
		log.Println("personID was not provided")
		NewError(http.StatusBadRequest, "personID is required").WriteError(w)
		return
	}

	var person Person
	err := json.NewDecoder(req.Body).Decode(&person)
	if err != nil {
		log.Println(fmt.Errorf("unable to decode Person. Error: %v", err))
		NewError(http.StatusBadRequest, "unable to unmarshal request").WriteError(w)
		return
	}

	id, err := uuid.FromString(pID)
	if err != nil {
		log.Println("id was not a uuid. ID:", pID)
		NewError(http.StatusBadRequest, "id was not a uuid").WriteError(w)
		return
	}

	p := &rent.Person{
		ID:           id,
		FirstName:    person.FirstName,
		LastName:     person.LastName,
		EmailAddress: person.EmailAddress,
	}

	err = r.personService.UpdatePerson(p)
	if err != nil {
		log.Println(fmt.Errorf("unable to update Person. PersonID: %s Error: %v", p.ID, err))
		NewError(http.StatusInternalServerError, "unable to update Person").WriteError(w)
		return
	}
}

func (r *Router) DeletePerson(w http.ResponseWriter, req *http.Request) {
	pID := chi.URLParam(req, personID)
	if pID == "" {
		log.Println("personID was not provided")
		NewError(http.StatusBadRequest, "personID is required").WriteError(w)
		return
	}

	id, err := uuid.FromString(pID)
	if err != nil {
		log.Println("id was not a uuid. ID:", pID)
		NewError(http.StatusBadRequest, "id was not a uuid").WriteError(w)
		return
	}

	err = r.personService.DeletePerson(id)
	if err != nil {
		log.Println(fmt.Errorf("unable to delete Person. PersonID: %s Error: %v", pID, err))
		NewError(http.StatusInternalServerError, "unable to delete Person").WriteError(w)
		return
	}
}
