package rent

import (
	uuid "github.com/satori/go.uuid"
)

type LandlordService interface {
	GetLandlord(landlordID uuid.UUID) (*Landlord, error)
	RegisterLandlord(landlord *Landlord) error
	CancelLandlord(landlordID uuid.UUID) error
}

type Landlord struct {
	person
	LandlordID uuid.UUID
}

func NewEmptyLandlord(username, password, firstName, lastName, emailAddress, phoneNumber string) (*Landlord, error) {
	p, err := newPerson(username, password, firstName, lastName, emailAddress, phoneNumber)
	if err != nil {
		return nil, err
	}
	return &Landlord{
		person:     *p,
		LandlordID: uuid.NewV4(),
	}, nil
}
