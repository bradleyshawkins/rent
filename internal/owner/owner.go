package owner

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type Owner struct {
	personID   uuid.UUID
	properties map[uuid.UUID]*Property
	address    Address
}

func NewOwner(a Address) (*Owner, error) {
	if err := a.validate(); err != nil {
		return nil, err
	}

	return &Owner{
		personID:   uuid.NewV4(),
		properties: make(map[uuid.UUID]*Property),
		address:    a,
	}, nil
}

func (o *Owner) AddProperty(p Property) error {
	if _, ok := o.properties[p.id]; ok {
		return fmt.Errorf("owner already owns property %s", p.id)
	}

	o.properties[p.id] = &p
	return nil
}
