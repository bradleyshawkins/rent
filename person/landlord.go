package person

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func (p *Service) CreateLandlord(personID uuid.UUID) (uuid.UUID, error) {
	landlordID, err := p.landlordDatastore.CreateLandlord(personID)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("unable to create landlord. Error: %w", err)
	}

	return landlordID, nil
}
