package person

import (
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/go-chi/chi"

	"github.com/bradleyshawkins/rent/rest"
)

func (p *Router) Cancel(w http.ResponseWriter, r *http.Request) {
	err := p.cancel(w, r)
	if err != nil {
		err.WriteError(w)
		return
	}
}

func (p *Router) cancel(w http.ResponseWriter, r *http.Request) *rest.Error {
	id := chi.URLParam(r, personID)
	if id == "" {
		return rest.NewError(http.StatusBadRequest, "personID is required")
	}

	personID, err := uuid.FromString(id)
	if err != nil {
		log.Println("invalid uuid received. UUID:", id)
		return rest.NewError(http.StatusBadRequest, "invalid personID provided")
	}

	err = p.personService.CancelPerson(personID)
	if err != nil {
		log.Println(fmt.Errorf("unable to cancel person. Error: %v, PersonID: %v", err, id))
		return rest.NewError(http.StatusInternalServerError, "unable to cancel person")
	}
	return nil
}
