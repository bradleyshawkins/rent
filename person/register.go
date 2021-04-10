package person

import (
	"fmt"

	uuid "github.com/satori/go.uuid"

	"github.com/bradleyshawkins/rent"
)

func (p *Service) Register(person *rent.Person) (uuid.UUID, error) {
	person.ID = uuid.NewV4()
	err := p.personDatastore.RegisterPerson(person)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("unable to register person. Error: %v", err)
	}
	return person.ID, nil
}

func (p *Service) GetPerson(id uuid.UUID) (*rent.Person, error) {
	person, err := p.personDatastore.GetPerson(id)
	if err != nil {
		return nil, fmt.Errorf("unable to get person. Error: %v", err)
	}
	return person, nil
}

func (p *Service) UpdatePerson(person *rent.Person) error {
	return nil
}
func (p *Service) DeletePerson(id uuid.UUID) error {
	return nil
}
