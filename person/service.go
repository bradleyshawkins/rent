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
	CreateLandlord(personID uuid.UUID) (uuid.UUID, error)
}

type Service struct {
	personDatastore   personDatastore
	landlordDatastore landlordDatastore
}

func NewPersonService(p personDatastore) *Service {
	return &Service{personDatastore: p}
}
