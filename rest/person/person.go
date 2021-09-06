package person

import (
	"log"

	"github.com/bradleyshawkins/rent"

	"github.com/bradleyshawkins/rent/rest"
	"github.com/go-chi/chi"
)

const (
	personID  = "personID"
	accountID = "accountID"
)

type Router struct {
	ps rent.PersonService
}

func NewRouter(ps rent.PersonService) *Router {
	return &Router{ps: ps}
}

func (p *Router) RegisterEndpoints(m chi.Router) {
	log.Println("Registering person endpoints")

	// Person management
	m.Post("/person/register", rest.ErrorHandler(p.RegisterPerson))

}
