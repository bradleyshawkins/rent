package landlord

import (
	"log"

	"github.com/bradleyshawkins/rent"

	"github.com/go-chi/chi"
)

const landlordID = "landlordID"

type Router struct {
	landlordService rent.LandlordService
}

func NewLandlordRouter(r rent.LandlordService) *Router {
	return &Router{landlordService: r}
}

func (p *Router) Register(m chi.Router) {
	log.Println("Registering landlord endpoints")

	m.Post("/landlord", p.RegisterLandlord)
	m.Delete("/landlord/{landlordID}", p.Cancel)
	m.Get("/landlord/{landlordID", p.GetLandlord)
}
