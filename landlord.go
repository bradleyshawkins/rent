package rent

import (
	uuid "github.com/satori/go.uuid"
)

type Landlord struct {
	Person
	LandlordID uuid.UUID
	Properties Properties
}

func NewEmptyLandlord(username, password, firstName, lastName, emailAddress, phoneNumber string) (*Landlord, error) {
	p, err := NewPerson(username, password, firstName, lastName, emailAddress, phoneNumber)
	if err != nil {
		return nil, err
	}
	return &Landlord{
		Person:     *p,
		LandlordID: uuid.NewV4(),
		Properties: NewEmptyProperties(),
	}, nil
}
