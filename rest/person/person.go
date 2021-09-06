package person

import (
	"log"

	"github.com/bradleyshawkins/rent"

	"github.com/bradleyshawkins/rent/rest"
	"github.com/go-chi/chi"
)

const (
	personID = "personID"
)

type Router struct {
	ps rent.PersonStore
}

func NewRouter(ps rent.PersonStore) *Router {
	return &Router{ps: ps}
}

func (p *Router) RegisterEndpoints(m chi.Router) {
	log.Println("Registering person endpoints")

	// Person management
	m.Post("/person/register", rest.ErrorHandler(p.RegisterPerson))
	m.Get("/person/{personID}", rest.ErrorHandler(p.LoadPerson))

}
