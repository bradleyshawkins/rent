package rent_test

import (
	"errors"
	"testing"

	"github.com/bradleyshawkins/rent"
	"github.com/matryer/is"
)

func TestNewPerson(t *testing.T) {
	is := is.New(t)

	u := "username"
	p := "password"
	fn := "firstName"
	ln := "lastName"
	ea := "test.email@test.com"
	pn := "8019991234"

	l, err := rent.NewPerson(u, p, fn, ln, ea, pn)
	is.NoErr(err)

	is.Equal(l.Username, u)
	is.Equal(l.Password, p)
	is.Equal(l.FirstName, fn)
	is.Equal(l.LastName, ln)
	is.Equal(l.EmailAddress, ea)
	is.Equal(l.PhoneNumber, pn)
}

func TestNewPerson_MissingField(t *testing.T) {
	u := "username"
	p := "password"
	fn := "firstName"
	ln := "lastName"
	ea := "test.email@test.com"
	pn := "8019991234"

	tests := []struct {
		name         string
		username     string
		password     string
		firstName    string
		lastName     string
		emailAddress string
		phoneNumber  string
	}{
		{name: "Missing username", password: p, firstName: fn, lastName: ln, emailAddress: ea, phoneNumber: pn},
		{name: "Missing password", username: u, firstName: fn, lastName: ln, emailAddress: ea, phoneNumber: pn},
		{name: "Missing firstName", username: u, password: p, lastName: ln, emailAddress: ea, phoneNumber: pn},
		{name: "Missing lastName", username: u, password: p, firstName: fn, emailAddress: ea, phoneNumber: pn},
		{name: "Missing emailAddress", username: u, password: p, firstName: fn, lastName: ln, phoneNumber: pn},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := is.New(t)
			_, err := rent.NewPerson(tt.username, tt.password, tt.firstName, tt.lastName, tt.emailAddress, tt.phoneNumber)
			i.True(err != nil)

			v := &rent.ValidationError{}
			i.True(errors.As(err, &v))
			i.True(v.Reason == rent.Missing)
		})
	}
}

func TestNewPerson_InvalidField(t *testing.T) {
	i := is.New(t)
	u := "username"
	p := "password"
	fn := "firstName"
	ln := "lastName"
	pn := "8019991234"

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
			_, err := rent.NewPerson(u, p, fn, ln, tt.emailAddress, pn)
			i.True(err != nil)

			v := &rent.ValidationError{}
			i.True(errors.As(err, &v))
			i.True(v.Reason == rent.Invalid)
		})
	}
}
