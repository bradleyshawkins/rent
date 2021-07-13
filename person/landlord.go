package person

import (
	"fmt"

	uuid "github.com/satori/go.uuid"

	"github.com/bradleyshawkins/rent"
)

func (p *Service) RegisterLandlord(landlord *rent.Landlord) error {
	err := p.landlordDatastore.RegisterLandlord(landlord)
	if err != nil {
		return fmt.Errorf("unable to create landlord. Error: %w", err)
	}

	return nil
}

func (p *Service) CancelLandlord(landlordID uuid.UUID) error {
	err := p.landlordDatastore.CancelLandlord(landlordID)
	if err != nil {
		return fmt.Errorf("unable to cancel landlord. Error: %w", err)
	}
	return nil
}
