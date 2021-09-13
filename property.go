package rent

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type PropertyStore interface {
	RegisterProperty(accountID uuid.UUID, p *Property) error
	LoadProperty(accountID uuid.UUID, propertyID uuid.UUID) (*Property, error)
	RemoveProperty(accountID uuid.UUID, propertyID uuid.UUID) error
}

type Property struct {
	ID      uuid.UUID
	Name    string
	Status  PropertyStatus
	Address *Address
}

type Address struct {
	ID      uuid.UUID
	Street1 string
	Street2 string
	City    string
	State   string
	Zipcode string
}

type PropertyStatus int

const (
	PropertyDisabled PropertyStatus = iota + 1
	PropertyVacant
	PropertyOccupied
)

var propertyStatusMap = map[PropertyStatus]string{
	PropertyDisabled: "Disabled",
	PropertyVacant:   "Vacant",
	PropertyOccupied: "Occupied",
}

func NewProperty(name string, address *Address) (*Property, error) {
	p := &Property{
		ID:      uuid.NewV4(),
		Name:    name,
		Status:  PropertyVacant,
		Address: address,
	}
	return p, p.Validate()
}

func NewExistingProperty(id uuid.UUID, name string, status PropertyStatus, address *Address) (*Property, error) {
	p := &Property{
		ID:      id,
		Name:    name,
		Status:  status,
		Address: address,
	}

	return p, p.Validate()
}

func (p *Property) Validate() error {
	if p.ID == (uuid.UUID{}) {
		return NewError(errors.New("property must have an ID"), WithInvalidFields(InvalidField{Field: "ID", Reason: ReasonMissing}))
	}
	if p.Name == "" {
		return NewError(errors.New("property must have a name"), WithInvalidFields(InvalidField{Field: "name", Reason: ReasonMissing}))
	}
	if p.Status == 0 {
		return NewError(errors.New("property must have a status"), WithInvalidFields(InvalidField{Field: "status", Reason: ReasonMissing}))
	}
	if _, ok := propertyStatusMap[p.Status]; !ok {
		return NewError(errors.New("invalid property status"), WithInvalidFields(InvalidField{Field: "status", Reason: ReasonInvalid}))
	}
	if err := p.Address.Validate(); err != nil {
		return err
	}
	return nil
}

func NewAddress(st1, st2, c, st, z string) (*Address, error) {
	a := &Address{
		ID:      uuid.NewV4(),
		Street1: st1,
		Street2: st2,
		City:    c,
		State:   st,
		Zipcode: z,
	}
	return a, a.Validate()
}

func NewExistingAddress(id uuid.UUID, st1, st2, c, st, z string) (*Address, error) {
	a := &Address{
		ID:      id,
		Street1: st1,
		Street2: st2,
		City:    c,
		State:   st,
		Zipcode: z,
	}

	return a, a.Validate()
}

func (a *Address) Validate() error {
	if a.ID == (uuid.UUID{}) {
		return NewError(errors.New("address must have an id"), WithInvalidFields(InvalidField{Field: "id", Reason: ReasonMissing}))
	}

	if a.Street1 == "" {
		return NewError(errors.New("address must have street1"), WithInvalidFields(InvalidField{Field: "street1", Reason: ReasonMissing}))
	}
	if a.City == "" {
		return NewError(errors.New("address must have city"), WithInvalidFields(InvalidField{Field: "city", Reason: ReasonMissing}))
	}
	if a.State == "" {
		return NewError(errors.New("address must have state"), WithInvalidFields(InvalidField{Field: "state", Reason: ReasonMissing}))
	}
	if a.Zipcode == "" {
		return NewError(errors.New("address must have zipcode"), WithInvalidFields(InvalidField{Field: "zipcode", Reason: ReasonMissing}))
	}
	return nil
}

func (p *Property) IsActive() error {
	if p.Status == PropertyDisabled {
		return NewError(errors.New("property is not active"), WithEntityDisabled())
	}
	return nil
}

func (p *Property) Disable() error {
	p.Status = PropertyDisabled
	return nil
}
