package identity_test

import (
	"errors"
	"github.com/bradleyshawkins/rent/identity"
	"github.com/bxcodec/faker/v3"
	"github.com/matryer/is"
	"net/mail"
	"testing"
)

type mockUserLoader struct {
	loadUserUserID identity.UserID
	loadUserRetVal *identity.User
	loadUserError  error
}

func (m *mockUserLoader) LoadUser(uID identity.UserID) (*identity.User, error) {
	m.loadUserUserID = uID
	return m.loadUserRetVal, m.loadUserError
}

func TestLoadUser(t *testing.T) {
	i := is.New(t)
	ea, err := mail.ParseAddress(faker.Email())
	i.NoErr(err)

	expected := &identity.User{
		ID:           identity.NewUserID(),
		EmailAddress: ea,
		FirstName:    faker.FirstName(),
		LastName:     faker.LastName(),
		Status:       identity.UserActive,
	}

	m := &mockUserLoader{
		loadUserRetVal: expected,
		loadUserError:  nil,
	}
	ul := identity.NewUserRetriever(m)

	actual, err := ul.LoadUser(expected.ID)
	i.NoErr(err)

	i.Equal(expected.ID, m.loadUserUserID)
	i.Equal(expected, actual)
}

func TestLoadUser_Error(t *testing.T) {
	i := is.New(t)

	m := &mockUserLoader{
		loadUserRetVal: nil,
		loadUserError:  errors.New("error loading user"),
	}
	ul := identity.NewUserRetriever(m)

	_, err := ul.LoadUser(identity.NewUserID())
	i.True(err != nil)
}