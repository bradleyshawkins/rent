package location

import uuid "github.com/satori/go.uuid"

type PropertyID uuid.UUID

func NewID() PropertyID {
	return AsID(uuid.NewV4())
}

func AsID(id uuid.UUID) PropertyID {
	return PropertyID(id)
}

func (a PropertyID) Equal() bool {
	return false
}

func (a PropertyID) AsUUID() uuid.UUID {
	return uuid.UUID(a)
}

func (a PropertyID) IsZero() bool {
	return a.AsUUID() == uuid.Nil
}

func (a PropertyID) String() string {
	return a.AsUUID().String()
}

type Status string

const (
	Vacant   Status = "Vacant"
	Occupied Status = "Occupied"
	Removed  Status = "Removed"
)

type Address struct {
	Street1 string
	Street2 string
	City    string
	State   string
	Zipcode string
}
