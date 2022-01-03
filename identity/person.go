package identity

import (
	uuid "github.com/satori/go.uuid"
)

type PersonID uuid.UUID

func NewPersonID() PersonID {
	return PersonID(uuid.NewV4())
}

func AsPersonID(id uuid.UUID) PersonID {
	return PersonID(id)
}

func (p PersonID) IsZero() bool {
	return p.AsUUID() == uuid.Nil
}

func (p PersonID) AsUUID() uuid.UUID {
	return uuid.UUID(p)
}

func (p PersonID) String() string {
	return p.AsUUID().String()
}

type PersonStatus string

const (
	PersonDisabled PersonStatus = "Disabled"
	PersonActive   PersonStatus = "Active"
)
