package account

import (
	"testing"

	"github.com/bradleyshawkins/rent"

	uuid "github.com/satori/go.uuid"
)

func TestRegisterAccount(t *testing.T) {

}

type mock struct {
	RegisterAccountAccountParam *rent.Account
	RegisterAccountError        error

	AddToAccountPersonParam    *rent.Person
	AddToAccountAccountIDParam uuid.UUID
	AddToAccountReturnError    error

	LoadAccountAccountIDParam uuid.UUID
	LoadAccountReturnAccount  *rent.Account
	LoadAccountReturnError    error

	RegisterPersonPersonParam *rent.Person
	RegisterPersonReturnError error
}

func (p *mock) LoadAccount(accountID uuid.UUID) (*rent.Account, error) {
	p.LoadAccountAccountIDParam = accountID
	return p.LoadAccountReturnAccount, p.LoadAccountReturnError
}

func (p *mock) RegisterAccount(a *rent.Account) error {
	p.RegisterAccountAccountParam = a
	return p.RegisterAccountError
}

func (p *mock) AddToAccount(aID uuid.UUID, pe *rent.Person) error {
	p.AddToAccountAccountIDParam = aID
	p.AddToAccountPersonParam = pe
	return p.AddToAccountReturnError
}

func (p *mock) RegisterPerson(person *rent.Person) error {
	p.RegisterPersonPersonParam = person
	return p.RegisterPersonReturnError
}
