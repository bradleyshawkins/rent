package rent_test

import (
	"testing"

	"github.com/bradleyshawkins/rent"

	"github.com/matryer/is"
)

func TestNewPerson(t *testing.T) {
	i := is.New(t)

	p := "password"
	fn := "firstName"
	ln := "lastName"
	ea := "test.email@test.com"

	l, err := rent.NewPerson(ea, p, fn, ln)
	i.NoErr(err)

	i.Equal(l.Password, p)
	i.Equal(l.FirstName, fn)
	i.Equal(l.LastName, ln)
	i.Equal(l.EmailAddress, ea)
	i.Equal(l.Status, rent.PersonActive)
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
			_, err := rent.NewPerson(tt.emailAddress, tt.password, tt.firstName, tt.lastName)
			i.True(err != nil)

			e, ok := err.(*rent.Error)
			i.True(ok)

			i.True(len(e.InvalidFields()) == 1)

			i.True(e.InvalidFields()[0].Reason == rent.ReasonMissing)
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
			_, err := rent.NewPerson(tt.emailAddress, p, fn, ln)
			i.True(err != nil)

			e, ok := err.(*rent.Error)
			i.True(ok)

			i.True(len(e.InvalidFields()) == 1)

			i.True(e.InvalidFields()[0].Reason == rent.ReasonInvalid)
		})
	}
}
