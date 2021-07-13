package person

import (
	"log"

	"github.com/bradleyshawkins/rent"
	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
)

const personID = "personID"

type personService interface {
	GetPerson(id uuid.UUID) (*rent.Person, error)
	CreatePerson(p *rent.Person) (uuid.UUID, error)
	UpdatePerson(p *rent.Person) error
	DeletePerson(id uuid.UUID) error
}

type Router struct {
	personService personService
}

func NewPersonRouter(u personService) *Router {
	return &Router{personService: u}
}

func (p *Router) Register(m chi.Router) {
	log.Println("Registering person endpoints")
	m.Post("/user", p.RegisterPerson)
	m.Get("/user/{userID}", p.GetPerson)
	m.Put("/user/{userID}", p.UpdatePerson)
	m.Delete("/user/{userID}", p.DeletePerson)
}
