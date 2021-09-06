package account

import (
	"github.com/bradleyshawkins/rent"
	uuid "github.com/satori/go.uuid"
)

type accountService interface {
	LoadAccount(id uuid.UUID) (*rent.Account, error)
	RegisterAccount(a *rent.Account) error
	AddToAccount(aID uuid.UUID, p *rent.Person) error
	RegisterPerson(p *rent.Person) error
}

type Service struct {
	as accountService
}

func NewService(as accountService) *Service {
	return &Service{as: as}
}

func (s *Service) RegisterAccount(emailAddress, password, firstName, lastName string) (uuid.UUID, error) {
	a, err := rent.NewAccount(s.as)
	if err != nil {
		return uuid.UUID{}, err
	}

	p, err := rent.NewPerson(s.as, emailAddress, password, firstName, lastName)
	if err != nil {
		return uuid.UUID{}, err
	}

	err = a.Register()
	if err != nil {
		return uuid.UUID{}, err
	}

	err = p.Register()
	if err != nil {
		return uuid.UUID{}, err
	}

	err = a.AddPerson(p)
	if err != nil {
		return uuid.UUID{}, err
	}

	return a.ID, nil
}

func (s *Service) AddPersonToAccount(accountID uuid.UUID, emailAddress, password, firstName, lastName string) error {
	p, err := rent.NewPerson(s.as, emailAddress, password, firstName, lastName)
	if err != nil {
		return err
	}

	err = p.Register()
	if err != nil {
		return err
	}

	a, err := s.as.LoadAccount(accountID)
	if err != nil {
		return err
	}

	err = a.AddPerson(p)
	if err != nil {
		return err
	}
	return nil
}
