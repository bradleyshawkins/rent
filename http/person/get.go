package person

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	h "github.com/bradleyshawkins/rent/http"
	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
)

type GetPersonResponse struct {
	ID           string `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
	PhoneNumber  string `json:"phoneNumber"`
}

func (r *Router) GetPerson(w http.ResponseWriter, req *http.Request) {
	if err := r.getPerson(w, req); err != nil {
		err.WriteError(w)
	}
}

func (r *Router) getPerson(w http.ResponseWriter, req *http.Request) *h.Error {
	log.Println("Starting get person request")
	pID := chi.URLParam(req, personID)
	if pID == "" {
		log.Println("personID was not provided")
		return h.NewError(http.StatusBadRequest, "userID is required")
	}

	id, err := uuid.FromString(pID)
	if err != nil {
		log.Println("id was not a uuid. ID:", pID)
		return h.NewError(http.StatusBadRequest, "id was not a uuid")
	}

	u, err := r.personService.GetPerson(id)
	if err != nil {
		log.Println(fmt.Errorf("unable to get person by id: %s Error: %v", pID, err))
		return h.NewError(http.StatusInternalServerError, "unable to get user by ID")
	}

	if u == nil {
		log.Println(fmt.Sprintf("person %s not found", id))
		return h.NewError(http.StatusNotFound, "user not found")
	}

	ur := GetPersonResponse{
		ID:           u.ID.String(),
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		EmailAddress: u.EmailAddress,
		PhoneNumber:  u.PhoneNumber,
	}

	err = json.NewEncoder(w).Encode(ur)
	if err != nil {
		log.Println(fmt.Errorf("unable to marshal person personID: %s, Error: %v", pID, err))
		return h.NewError(http.StatusInternalServerError, "unable to marshal user")
	}

	return nil
}
