package account

import (
	"log"

	"github.com/bradleyshawkins/rent/account"

	"github.com/go-chi/chi"
)

const (
	personID  = "personID"
	accountID = "accountID"
)

type Router struct {
	as *account.Service
}

func NewRouter(as *account.Service) *Router {
	return &Router{as: as}
}

func (p *Router) RegisterEndpoints(m chi.Router) {
	log.Println("Registering person endpoints")

	// Account Management
	m.Post("/accounts/register", p.RegisterAccount)
	m.Post("/accounts/{accountID}/register", p.RegisterPersonToAccount)
	m.Get("/accounts/{accountID}/settings", nil)
	m.Put("/accounts/{accountID}/settings", nil)

	// Person management
	m.Post("/person/register", p.RegisterAccount)
	m.Delete("/person/{personID}/cancel", p.Cancel)
	m.Put("/person/{personID}/password", nil)

	// Person Settings
	m.Get("/person/{personID}/settings", p.Get)
	m.Put("/person/{personID}/settings", nil)
}
