package person

import (
	"log"

	"github.com/bradleyshawkins/rent"

	"github.com/go-chi/chi"
)

const personID = "personID"

type Router struct {
	personService rent.PersonService
}

func NewPersonRouter(r rent.PersonService) *Router {
	return &Router{personService: r}
}

func (p *Router) RegisterEndpoints(m chi.Router) {
	log.Println("Registering person endpoints")

	m.Post("/person", p.Register)
	m.Delete("/person/{personID}", p.Cancel)
	m.Get("/person/{personID}", p.Get)
}
