package person

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bradleyshawkins/rent"
	h "github.com/bradleyshawkins/rent/http"
	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
)

type UpdatePersonRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"` // TODO: Use better password handling
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
	PhoneNumber  string `json:"PhoneNumber"`
}

func (r *Router) UpdatePerson(w http.ResponseWriter, req *http.Request) {
	if err := r.updatePerson(w, req); err != nil {
		err.WriteError(w)
	}
}

func (r *Router) updatePerson(w http.ResponseWriter, req *http.Request) *h.Error {
	pID := chi.URLParam(req, personID)
	if pID == "" {
		log.Println("userID was not provided")
		return h.NewError(http.StatusBadRequest, "userID is required")
	}

	var rr UpdatePersonRequest
	err := json.NewDecoder(req.Body).Decode(&rr)
	if err != nil {
		log.Println(fmt.Errorf("unable to decode user. Error: %v", err))
		return h.NewError(http.StatusBadRequest, "unable to unmarshal request")
	}

	id, err := uuid.FromString(pID)
	if err != nil {
		log.Println("id was not a uuid. ID:", pID)
		return h.NewError(http.StatusBadRequest, "id was not a uuid")
	}

	user, err := rent.NewPersonWithID(id, rr.Username, rr.Password, rr.FirstName, rr.LastName, rr.EmailAddress, rr.PhoneNumber)
	if err != nil {
		log.Println(fmt.Errorf("error creating person. Error: %v", err))
		return h.NewError(http.StatusBadRequest, "missing field")
	}
	err = r.personService.UpdatePerson(user)
	if err != nil {
		log.Println(fmt.Errorf("unable to update user. UserID: %s Error: %v", user.ID, err))
		return h.NewError(http.StatusInternalServerError, "unable to update user")
	}

	return nil
}
