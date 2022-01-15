package identity_test

import (
	"errors"
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/matryer/is"

	"github.com/bradleyshawkins/rent/identity"
)

type mockUserCreatorService struct {
	RegisterUserUserRegistration       *identity.UserRegistration
	RegisterUserError                  error
	RegisterAccountUserID              identity.UserID
	RegisterAccountAccountRegistration *identity.AccountRegistration
	RegisterAccountError               error
	AddUserToAccountAccountID          identity.AccountID
	AddUserToAccountUserID             identity.UserID
	AddUserToAccountRole               identity.Role
	AddUserToAccountError              error
}

func (m *mockUserCreatorService) RegisterUser(p *identity.UserRegistration) error {
	m.RegisterUserUserRegistration = p
	return m.RegisterUserError
}

func (m *mockUserCreatorService) RegisterAccount(pID identity.UserID, a *identity.AccountRegistration) error {
	m.RegisterAccountUserID = pID
	m.RegisterAccountAccountRegistration = a
	return m.RegisterAccountError
}

func (m *mockUserCreatorService) AddUserToAccount(aID identity.AccountID, pID identity.UserID, role identity.Role) error {
	m.AddUserToAccountAccountID = aID
	m.AddUserToAccountUserID = pID
	m.AddUserToAccountRole = role
	return m.AddUserToAccountError
}

type mockUserCreator struct {
	mpcs *mockUserCreatorService
}

func (m *mockUserCreator) Register(f identity.RegistrationFunc) error {
	return f(m.mpcs)
}

func TestRegisterUser(t *testing.T) {
	i := is.New(t)
	mpcs := &mockUserCreatorService{}
	mpc := &mockUserCreator{mpcs}
	registrar := identity.NewRegistrar(mpc)
	emailAddress := "email.address@test.com"
	firstName := "First"
	lastName := "Last"
	password := "Password"

	user, account, err := registrar.Register(emailAddress, firstName, lastName, password)
	if err != nil {
		t.Fatal("Unexpected error occurred. Error:", err)
	}

	if user == nil {
		t.Fatal("User was nil. User:", user)
	}

	t.Log("User:", user)

	i.Equal(user.EmailAddress.Address, emailAddress)
	i.Equal(user.FirstName, firstName)
	i.Equal(user.LastName, lastName)
	i.True(!user.ID.IsZero())

	if account == nil {
		t.Fatal("Account was nil. Account:", account)
	}

	t.Log("Account:", account)

	i.True(!account.ID.IsZero())
}

func TestRegisterUser_Fail(t *testing.T) {
	tests := []struct {
		name               string
		createUserError    error
		createAccountError error
	}{
		{
			name:               "Create User Fails",
			createUserError:    errors.New("unable to create user"),
			createAccountError: nil,
		},
		{
			name:               "Create Account Fails",
			createUserError:    nil,
			createAccountError: errors.New("unable to create account"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mpcs := &mockUserCreatorService{
				RegisterUserError:    tt.createUserError,
				RegisterAccountError: tt.createAccountError,
			}
			mpc := &mockUserCreator{mpcs: mpcs}

			registrar := identity.NewRegistrar(mpc)

			user, account, err := registrar.Register("email.address@test.com", "First", "Last", "Password")
			if err == nil {
				t.Error("Received a nil error. Should have received an error")
			}
			if user != nil {
				t.Error("user was non-nil. User:", user)
			}
			if account != nil {
				t.Error("account was non-nil. Account:", account)
			}
			t.Logf("error: %v", err)
		})
	}
}

func TestRegisterUserToAccount(t *testing.T) {
	i := is.New(t)
	mpcs := &mockUserCreatorService{}
	mpc := &mockUserCreator{mpcs}
	registrar := identity.NewRegistrar(mpc)
	emailAddress := "email.address@test.com"
	firstName := "First"
	lastName := "Last"
	password := "Password"
	accountID := identity.AsAccountID(uuid.NewV4())

	user, err := registrar.RegisterUserToAccount(accountID, "Owner", emailAddress, firstName, lastName, password)
	if err != nil {
		t.Fatal("Unexpected error occurred. Error:", err)
	}

	if user == nil {
		t.Fatal("User was nil. User:", user)
	}

	t.Log("User:", user)

	i.Equal(user.EmailAddress.Address, emailAddress)
	i.Equal(user.FirstName, firstName)
	i.Equal(user.LastName, lastName)
	i.True(!user.ID.IsZero())

	i.Equal(mpcs.AddUserToAccountRole, identity.RoleOwner)
	i.Equal(mpcs.AddUserToAccountAccountID, accountID)
	i.Equal(mpcs.AddUserToAccountUserID, user.ID)
}

func TestRegisterUserToAccount_Fail(t *testing.T) {
	tests := []struct {
		name                  string
		registerUserError     error
		addUserToAccountError error
	}{
		{
			name:                  "Create User Fails",
			registerUserError:     errors.New("unable to create user"),
			addUserToAccountError: nil,
		},
		{
			name:                  "Create Account Fails",
			registerUserError:     nil,
			addUserToAccountError: errors.New("unable to add user to account"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mpcs := &mockUserCreatorService{
				RegisterUserError:     tt.registerUserError,
				AddUserToAccountError: tt.addUserToAccountError,
			}
			mpc := &mockUserCreator{mpcs: mpcs}

			registrar := identity.NewRegistrar(mpc)

			user, err := registrar.RegisterUserToAccount(identity.NewAccountID(), "Owner", "email.address@test.com", "First", "Last", "Password")
			if err == nil {
				t.Error("Received a nil error. Should have received an error")
			}
			if user != nil {
				t.Error("user was non-nil. User:", user)
			}
			t.Logf("error: %v", err)
		})
	}
}
