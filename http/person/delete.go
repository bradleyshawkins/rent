package person

import (
	"fmt"
	"log"
	"net/http"

	h "github.com/bradleyshawkins/rent/http"
	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
)

func (r *Router) DeletePerson(w http.ResponseWriter, req *http.Request) {
	if err := r.deletePerson(w, req); err != nil {
		err.WriteError(w)
	}
}

func (r *Router) deletePerson(w http.ResponseWriter, req *http.Request) *h.Error {
	pID := chi.URLParam(req, personID)
	if pID == "" {
		log.Println("userID was not provided")
		return h.NewError(http.StatusBadRequest, "userID is required")
	}

	id, err := uuid.FromString(pID)
	if err != nil {
		log.Println("id was not a uuid. ID:", pID)
		return h.NewError(http.StatusBadRequest, "id was not a uuid")
	}

	err = r.personService.DeletePerson(id)
	if err != nil {
		log.Println(fmt.Errorf("unable to delete user. userID: %s Error: %v", pID, err))
		return h.NewError(http.StatusInternalServerError, "unable to delete user")
	}

	return nil
}
