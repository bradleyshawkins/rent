package identity_test

import (
	"errors"
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/matryer/is"

	"github.com/bradleyshawkins/rent/identity"
)

type mockPersonCreatorService struct {
	RegisterPersonPersonRegistration   *identity.PersonRegistration
	RegisterPersonError                error
	RegisterAccountPersonID            identity.PersonID
	RegisterAccountAccountRegistration *identity.AccountRegistration
	RegisterAccountError               error
	AddPersonToAccountAccountID        identity.AccountID
	AddPersonToAccountPersonID         identity.PersonID
	AddPersonToAccountRole             identity.Role
	AddPersonToAccountError            error
}

func (m *mockPersonCreatorService) RegisterPerson(p *identity.PersonRegistration) error {
	m.RegisterPersonPersonRegistration = p
	return m.RegisterPersonError
}

func (m *mockPersonCreatorService) RegisterAccount(pID identity.PersonID, a *identity.AccountRegistration) error {
	m.RegisterAccountPersonID = pID
	m.RegisterAccountAccountRegistration = a
	return m.RegisterAccountError
}

func (m *mockPersonCreatorService) AddPersonToAccount(aID identity.AccountID, pID identity.PersonID, role identity.Role) error {
	m.AddPersonToAccountAccountID = aID
	m.AddPersonToAccountPersonID = pID
	m.AddPersonToAccountRole = role
	return m.AddPersonToAccountError
}

type mockPersonCreator struct {
	mpcs *mockPersonCreatorService
}

func (m *mockPersonCreator) Register(f identity.RegistrationFunc) error {
	return f(m.mpcs)
}

func TestRegisterPerson(t *testing.T) {
	i := is.New(t)
	mpcs := &mockPersonCreatorService{}
	mpc := &mockPersonCreator{mpcs}
	registrar := identity.NewRegistrar(mpc)
	emailAddress := "email.address@test.com"
	firstName := "First"
	lastName := "Last"
	password := "Password"

	person, account, err := registrar.Register(emailAddress, firstName, lastName, password)
	if err != nil {
		t.Fatal("Unexpected error occurred. Error:", err)
	}

	if person == nil {
		t.Fatal("Person was nil. Person:", person)
	}

	t.Log("Person:", person)

	i.Equal(person.EmailAddress.Address, emailAddress)
	i.Equal(person.FirstName, firstName)
	i.Equal(person.LastName, lastName)
	i.True(!person.ID.IsZero())

	if account == nil {
		t.Fatal("Account was nil. Account:", account)
	}

	t.Log("Account:", account)

	i.True(!account.ID.IsZero())
}

func TestRegisterPerson_Fail(t *testing.T) {
	tests := []struct {
		name               string
		createPersonError  error
		createAccountError error
	}{
		{
			name:               "Create Person Fails",
			createPersonError:  errors.New("unable to create person"),
			createAccountError: nil,
		},
		{
			name:               "Create Account Fails",
			createPersonError:  nil,
			createAccountError: errors.New("unable to create account"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mpcs := &mockPersonCreatorService{
				RegisterPersonError:  tt.createPersonError,
				RegisterAccountError: tt.createAccountError,
			}
			mpc := &mockPersonCreator{mpcs: mpcs}

			registrar := identity.NewRegistrar(mpc)

			person, account, err := registrar.Register("email.address@test.com", "First", "Last", "Password")
			if err == nil {
				t.Error("Received a nil error. Should have received an error")
			}
			if person != nil {
				t.Error("person was non-nil. Person:", person)
			}
			if account != nil {
				t.Error("account was non-nil. Account:", account)
			}
			t.Logf("error: %v", err)
		})
	}
}

func TestRegisterPersonToAccount(t *testing.T) {
	i := is.New(t)
	mpcs := &mockPersonCreatorService{}
	mpc := &mockPersonCreator{mpcs}
	registrar := identity.NewRegistrar(mpc)
	emailAddress := "email.address@test.com"
	firstName := "First"
	lastName := "Last"
	password := "Password"
	accountID := identity.AsAccountID(uuid.NewV4())

	person, err := registrar.RegisterPersonToAccount(accountID, "Owner", emailAddress, firstName, lastName, password)
	if err != nil {
		t.Fatal("Unexpected error occurred. Error:", err)
	}

	if person == nil {
		t.Fatal("Person was nil. Person:", person)
	}

	t.Log("Person:", person)

	i.Equal(person.EmailAddress.Address, emailAddress)
	i.Equal(person.FirstName, firstName)
	i.Equal(person.LastName, lastName)
	i.True(!person.ID.IsZero())

	i.Equal(mpcs.AddPersonToAccountRole, identity.RoleOwner)
	i.Equal(mpcs.AddPersonToAccountAccountID, accountID)
	i.Equal(mpcs.AddPersonToAccountPersonID, person.ID)
}

func TestRegisterPersonToAccount_Fail(t *testing.T) {
	tests := []struct {
		name                    string
		registerPersonError     error
		addPersonToAccountError error
	}{
		{
			name:                    "Create Person Fails",
			registerPersonError:     errors.New("unable to create person"),
			addPersonToAccountError: nil,
		},
		{
			name:                    "Create Account Fails",
			registerPersonError:     nil,
			addPersonToAccountError: errors.New("unable to add person to account"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mpcs := &mockPersonCreatorService{
				RegisterPersonError:     tt.registerPersonError,
				AddPersonToAccountError: tt.addPersonToAccountError,
			}
			mpc := &mockPersonCreator{mpcs: mpcs}

			registrar := identity.NewRegistrar(mpc)

			person, err := registrar.RegisterPersonToAccount(identity.NewAccountID(), "Owner", "email.address@test.com", "First", "Last", "Password")
			if err == nil {
				t.Error("Received a nil error. Should have received an error")
			}
			if person != nil {
				t.Error("person was non-nil. Person:", person)
			}
			t.Logf("error: %v", err)
		})
	}
}
