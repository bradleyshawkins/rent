package rent

import (
	"testing"

	"github.com/matryer/is"
)

func TestNewAccount(t *testing.T) {
	i := is.New(t)
	a, err := NewAccount()
	i.NoErr(err)

	i.True(a != nil)
}

func TestNewAccount_Error(t *testing.T) {
	i := is.New(t)

	_, err := NewAccount()
	i.True(err != nil)
}

//
//func TestAccountRegister(t *testing.T) {
//	i := is.New(t)
//	m := &account.mock{}
//	a, err := NewAccount(m)
//	i.NoErr(err)
//
//	err = a.Register()
//	i.NoErr(err)
//
//	account := m.RegisterAccountAccountParam
//	i.Equal(account.ID, a.ID)
//	i.Equal(account.Status, AccountActive)
//}
//
//func TestAccountRegister_Error(t *testing.T) {
//	i := is.New(t)
//
//	m := &account.mock{
//		RegisterAccountError: errors.New("error registering account"),
//	}
//	a, err := NewAccount(m)
//	i.NoErr(err)
//
//	err = a.Register()
//	i.True(err != nil)
//}
//
//func TestAddPerson(t *testing.T) {
//	i := is.New(t)
//	m := &account.mock{}
//	a, err := NewAccount(m)
//	i.NoErr(err)
//
//	p, err := NewPerson(m, "test.tester@test.com", "password", "firstName", "lastName")
//	i.NoErr(err)
//
//	err = a.AddPerson(p)
//	i.NoErr(err)
//
//	person := m.AddToAccountPersonParam
//	i.Equal(m.AddToAccountAccountIDParam, a.ID)
//	i.Equal(*p, *person)
//}
//
//func TestAddPerson_Error(t *testing.T) {
//	i := is.New(t)
//	m := &account.mock{
//		AddToAccountReturnError: errors.New("unable to add person to account"),
//	}
//	a, err := NewAccount(m)
//	i.NoErr(err)
//
//	p, err := NewPerson(m, "test.tester@test.com", "password", "firstName", "lastName")
//	i.NoErr(err)
//
//	err = a.AddPerson(p)
//	i.True(err != nil)
//}
