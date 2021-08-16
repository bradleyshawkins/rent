package account

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestRegisterAccount(t *testing.T) {

}

type mock struct {
	RegisterAccountAccountParam *Account
	RegisterAccountError        error

	AddToAccountPersonParam    *Person
	AddToAccountAccountIDParam uuid.UUID
	AddToAccountReturnError    error

	LoadAccountAccountIDParam uuid.UUID
	LoadAccountReturnAccount  *Account
	LoadAccountReturnError    error

	RegisterPersonPersonParam *Person
	RegisterPersonReturnError error
}

func (p *mock) LoadAccount(accountID uuid.UUID) (*Account, error) {
	p.LoadAccountAccountIDParam = accountID
	return p.LoadAccountReturnAccount, p.LoadAccountReturnError
}

func (p *mock) RegisterAccount(a *Account) error {
	p.RegisterAccountAccountParam = a
	return p.RegisterAccountError
}

func (p *mock) AddToAccount(aID uuid.UUID, pe *Person) error {
	p.AddToAccountAccountIDParam = aID
	p.AddToAccountPersonParam = pe
	return p.AddToAccountReturnError
}

func (p *mock) RegisterPerson(person *Person) error {
	p.RegisterPersonPersonParam = person
	return p.RegisterPersonReturnError
}
