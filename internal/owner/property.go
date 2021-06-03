package owner

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type PropertyStore interface {
	PropertyExists(id uuid.UUID) (bool, error)
}

type Property struct {
	id           uuid.UUID
	address      Address
	propertyInfo PropertyInfo
	renters      []uuid.UUID
}

type PropertyInfo struct {
	YearBuilt    time.Time
	HouseType    string
	RentPerMonth string
	PropertySize string
	LandSize     string
	Restrictions string // No Pets, No Smoking, etc
}

type Address struct {
	Street1 string
	Street2 string
	City    string
	State   string
	Zipcode string
}

func (a Address) validate() error {
	if a.Street1 == "" {
		return errors.New("street1 is required")
	}
	if a.City == "" {
		return errors.New("city is required")
	}
	if a.State == "" {
		return errors.New("state is required")
	}
	if a.Zipcode == "" {
		return errors.New("zipcode is required")
	}

	return nil
}

func NewProperty(a Address, p PropertyInfo) (*Property, error) {
	if err := a.validate(); err != nil {
		return nil, err
	}

	return &Property{
		id:           uuid.NewV4(),
		propertyInfo: p,
		address:      a,
	}, nil
}

func (p *Property) IsOccupied() bool {
	return len(p.renters) > 0
}
