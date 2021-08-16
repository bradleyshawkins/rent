package account

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type Account struct {
	ID     uuid.UUID
	Status AccountStatus

	as accountService
}

type AccountStatus int

const (
	AccountActive AccountStatus = iota + 1
	AccountInactive
	AccountDisabled
)

func NewAccount(as accountService) (*Account, error) {
	a := &Account{
		ID:     uuid.NewV4(),
		Status: AccountActive,
		as:     as,
	}
	return a, a.validate()
}

func NewAccountWithID(as accountService, id uuid.UUID, status AccountStatus) (*Account, error) {
	a := &Account{
		ID:     id,
		Status: status,
		as:     as,
	}
	return a, a.validate()
}

func (a *Account) validate() error {
	if a == nil {
		return errors.New("account has not been initialized")
	}

	if a.ID == (uuid.UUID{}) {
		return errors.New("accountID has not been set")
	}

	if a.as == nil {
		return errors.New("accountService is null")
	}

	return nil
}

func (a *Account) Register() error {
	if err := a.validate(); err != nil {
		return err
	}

	err := a.as.RegisterAccount(a)
	if err != nil {
		return err
	}
	return nil
}

func (a *Account) AddPerson(p *Person) error {
	if err := a.validate(); err != nil {
		return err
	}

	err := a.as.AddToAccount(a.ID, p)
	if err != nil {
		return err
	}
	return nil
}
