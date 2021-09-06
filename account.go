package rent

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type Account struct {
	ID     uuid.UUID
	Status AccountStatus
}

type AccountStatus int

const (
	AccountActive AccountStatus = iota + 1
	AccountInactive
	AccountDisabled
)

func NewAccount() (*Account, error) {
	a := &Account{
		ID:     uuid.NewV4(),
		Status: AccountActive,
	}
	return a, a.Validate()
}

func NewAccountWithID(id uuid.UUID, status AccountStatus) (*Account, error) {
	a := &Account{
		ID:     id,
		Status: status,
	}
	return a, a.Validate()
}

func (a *Account) Validate() error {
	if a == nil {
		return errors.New("account has not been initialized")
	}

	if a.ID == (uuid.UUID{}) {
		return errors.New("accountID has not been set")
	}

	return nil
}

//func (a *Account) Register() error {
//	if err := a.Validate(); err != nil {
//		return err
//	}
//
//	err := a.as.RegisterAccount(a)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (a *Account) AddPerson(p *Person) error {
//	if err := a.Validate(); err != nil {
//		return err
//	}
//
//	err := a.as.AddToAccount(a.ID, p)
//	if err != nil {
//		return err
//	}
//	return nil
//}
