package renter

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type Renter struct {
	personID uuid.UUID
	home     *Property
}

type Property struct {
	id      uuid.UUID
	ownerID uuid.UUID
}

func (r *Renter) MoveIn(p *Property) error {
	if r.home != nil {
		return errors.New("renter already lives in a home")
	}

	r.home = p
	return nil
}

func (r *Renter) MoveOut() error {
	if r.home == nil {
		return errors.New("renter does not live in a home")
	}

	r.home = nil
	return nil
}
