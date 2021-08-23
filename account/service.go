package account

import uuid "github.com/satori/go.uuid"

type accountService interface {
	LoadAccount(id uuid.UUID) (*Account, error)
	RegisterAccount(a *Account) error
	AddToAccount(aID uuid.UUID, p *Person) error
	RegisterPerson(p *Person) error
}

type Service struct {
	as accountService
}

func NewService(as accountService) *Service {
	return &Service{as: as}
}

func (s *Service) RegisterAccount(emailAddress, password, firstName, lastName string) (uuid.UUID, error) {
	a, err := NewAccount(s.as)
	if err != nil {
		return uuid.UUID{}, err
	}

	p, err := NewPerson(s.as, emailAddress, password, firstName, lastName)
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
	p, err := NewPerson(s.as, emailAddress, password, firstName, lastName)
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
