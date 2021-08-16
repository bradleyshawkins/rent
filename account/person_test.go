package account

import (
	"errors"
	"testing"

	"github.com/bradleyshawkins/rent/types"

	"github.com/matryer/is"
)

func TestNewPerson(t *testing.T) {
	i := is.New(t)

	p := "password"
	fn := "firstName"
	ln := "lastName"
	ea := "test.email@test.com"

	l, err := NewPerson(&mock{}, ea, p, fn, ln)
	i.NoErr(err)

	i.Equal(l.Password, p)
	i.Equal(l.FirstName, fn)
	i.Equal(l.LastName, ln)
	i.Equal(l.EmailAddress, ea)
	i.Equal(l.Status, PersonActive)
}

func TestNewPerson_MissingField(t *testing.T) {
	p := "password"
	fn := "firstName"
	ln := "lastName"
	ea := "test.email@test.com"

	tests := []struct {
		name         string
		password     string
		firstName    string
		lastName     string
		emailAddress string
	}{
		{name: "Missing password", firstName: fn, lastName: ln, emailAddress: ea},
		{name: "Missing firstName", password: p, lastName: ln, emailAddress: ea},
		{name: "Missing lastName", password: p, firstName: fn, emailAddress: ea},
		{name: "Missing emailAddress", password: p, firstName: fn, lastName: ln},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := is.New(t)
			_, err := NewPerson(&mock{}, tt.emailAddress, tt.password, tt.firstName, tt.lastName)
			i.True(err != nil)

			v := &types.FieldValidationError{}
			i.True(errors.As(err, &v))
			i.True(v.Reason == types.Missing)
		})
	}
}

func TestNewPerson_InvalidField(t *testing.T) {
	i := is.New(t)
	p := "password"
	fn := "firstName"
	ln := "lastName"

	tests := []struct {
		name         string
		emailAddress string
	}{
		{name: "no domain", emailAddress: "test.test"},
		{name: "domain only", emailAddress: "@test.com"},
		{name: "missing domain with signature", emailAddress: "Test Tester <email@example.com>"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewPerson(&mock{}, tt.emailAddress, p, fn, ln)
			i.True(err != nil)

			v := &types.FieldValidationError{}
			i.True(errors.As(err, &v))
			i.True(v.Reason == types.Invalid)
		})
	}
}

func TestRegister(t *testing.T) {
	i := is.New(t)

	ea := "test.tester@test.com"
	pass := "password"
	fn := "firstName"
	ln := "lastName"

	m := &mock{}
	p, err := NewPerson(m, ea, pass, fn, ln)
	i.NoErr(err)

	err = p.Register()
	i.NoErr(err)

	registeredPerson := m.RegisterPersonPersonParam

	i.Equal(registeredPerson.EmailAddress, ea)
	i.Equal(registeredPerson.Password, pass)
	i.Equal(registeredPerson.FirstName, fn)
	i.Equal(registeredPerson.LastName, ln)
	i.Equal(registeredPerson.Status, PersonActive)
}

func TestRegister_ReturnsError(t *testing.T) {
	i := is.New(t)

	m := &mock{
		RegisterPersonReturnError: errors.New("error registering person"),
	}

	p, err := NewPerson(m, "test.tester@test.com", "password", "firstName", "lastName")
	i.NoErr(err)

	err = p.Register()
	i.True(err != nil)
}
