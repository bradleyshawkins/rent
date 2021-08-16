package account

import (
	"net/http"

	"github.com/bradleyshawkins/rent/rest"
	uuid "github.com/satori/go.uuid"
)

type GetPersonResponse struct {
	PersonID  uuid.UUID
	FirstName string
	LastName  string
}

func (p *Router) Get(w http.ResponseWriter, r *http.Request) {
	err := p.get(w, r)
	if err != nil {
		err.WriteError(w)
		return
	}
}

func (p *Router) get(w http.ResponseWriter, r *http.Request) *rest.Error {
	//id := chi.URLParam(r, personID)
	//if id == "" {
	//	return rest.NewError(http.StatusBadRequest, "personID is required")
	//}
	//
	//personID, err := uuid.FromString(id)
	//if err != nil {
	//	log.Println("invalid uuid received. UUID:", id)
	//	return rest.NewError(http.StatusBadRequest, "invalid personID provided")
	//}
	//
	//l, err := p.personService.GetPerson(personID)
	//if err != nil {
	//	log.Printf("unable to get person. PersonID: %v Error: %v\n", id, err)
	//	return rest.NewError(http.StatusInternalServerError, "unable to retrieve person")
	//}
	//
	//err = json.NewEncoder(w).Encode(GetPersonResponse{
	//	PersonID:  l.ID,
	//	FirstName: l.FirstName,
	//	LastName:  l.LastName,
	//})
	//if err != nil {
	//	log.Printf("unable to encode person response. PersonID: %v, Error: %v\n", id, err)
	//	return rest.NewError(http.StatusInternalServerError, "unable to marshal person")
	//}
	return nil
}
