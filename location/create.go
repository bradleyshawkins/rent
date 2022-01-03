package location

import "github.com/bradleyshawkins/rent/identity"

type PropertyCreation struct {
	ID      PropertyID
	Name    string
	Status  Status
	Address *Address
}

type creator interface {
	CreateProperty(c CreateFunc) error
}

type PropertyCreatorService interface {
	CreateProperty(accountID identity.AccountID, c *PropertyCreation) error
}

type PropertyCreator struct {
	c creator
}

func NewCreator(c creator) *PropertyCreator {
	return &PropertyCreator{c: c}
}

// CreateFunc is a closer around the steps used to create a property
type CreateFunc func(us PropertyCreatorService) error

func (c *PropertyCreator) Create(accountID identity.AccountID, name string, a *Address) (*PropertyCreation, error) {
	creation := &PropertyCreation{
		ID:      NewID(),
		Name:    name,
		Status:  Vacant,
		Address: a,
	}

	err := c.c.CreateProperty(c.create(accountID, creation))
	if err != nil {
		return nil, err
	}
	return creation, nil
}

func (c *PropertyCreator) create(accountID identity.AccountID, pc *PropertyCreation) CreateFunc {
	return func(cs PropertyCreatorService) error {
		err := cs.CreateProperty(accountID, pc)
		if err != nil {
			return err
		}
		return nil
	}
}
