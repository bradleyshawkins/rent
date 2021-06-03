package landlord

import (
	"log"

	uuid "github.com/satori/go.uuid"

	"github.com/go-chi/chi"
)

const landlordID = "landlordID"

type landlordService interface {
	CreateLandlord(personID uuid.UUID) (uuid.UUID, error)
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
}
