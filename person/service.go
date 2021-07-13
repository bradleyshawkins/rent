package person

import (
	"github.com/bradleyshawkins/rent"
	uuid "github.com/satori/go.uuid"
)

type personDatastore interface {
	GetPerson(id uuid.UUID) (*rent.Person, error)
	CreatePerson(p *rent.Person) error
	UpdatePerson(p *rent.Person) error
	DeletePerson(id uuid.UUID) error
}

type landlordDatastore interface {
	RegisterLandlord(landlord *rent.Landlord) error
	CancelLandlord(landlordID uuid.UUID) error
}

type Service struct {
	personDatastore   personDatastore
	landlordDatastore landlordDatastore
}

func NewPersonService(p personDatastore, l landlordDatastore) *Service {
	return &Service{personDatastore: p, landlordDatastore: l}
}
