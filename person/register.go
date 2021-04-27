package person

import (
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"

	"github.com/bradleyshawkins/rent"
)

func (p *Service) Register(person *rent.Person) (uuid.UUID, error) {
	person.ID = uuid.NewV4()
	err := p.personDatastore.RegisterPerson(person)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("unable to register person. Error: %w", err)
	}
	return person.ID, nil
}

func (p *Service) GetPerson(id uuid.UUID) (*rent.Person, error) {
	log.Println("Getting person", id)
	person, err := p.personDatastore.GetPerson(id)
	if err != nil {
		return nil, fmt.Errorf("unable to get person. Error: %w", err)
	}

	log.Printf("Person: %+v\n", person)
	return person, nil
}

func (p *Service) UpdatePerson(person *rent.Person) error {
	err := p.personDatastore.UpdatePerson(person)
	if err != nil {
		return fmt.Errorf("unable to update person. Error: %w", err)
	}
	return nil
}

func (p *Service) DeletePerson(id uuid.UUID) error {
	err := p.personDatastore.DeletePerson(id)
	if err != nil {
		return fmt.Errorf("unable to delete person. Error: %w", err)
	}
	return nil
}
