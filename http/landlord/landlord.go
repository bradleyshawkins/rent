package landlord

import (
	"log"

	uuid "github.com/satori/go.uuid"

	"github.com/bradleyshawkins/rent"

	"github.com/go-chi/chi"
)

const landlordID = "landlordID"

type landlordService interface {
	RegisterLandlord(landlord *rent.Landlord) error
	CancelLandlord(landlordID uuid.UUID) error
	GetLandlord(landlordID uuid.UUID) (rent.Landlord, error)
}

type Router struct {
	landlordService landlordService
}

func NewLandlordRouter(r landlordService) *Router {
	return &Router{landlordService: r}
}

func (p *Router) Register(m chi.Router) {
	log.Println("Registering landlord endpoints")
	// Associate a person as a landlord
	m.Post("/landlord", p.RegisterLandlord)
	m.Delete("/landlord/{landlordID}", p.Cancel)
}
