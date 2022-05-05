package identity_test

import (
	"errors"
	"net/mail"
	"testing"

	identity2 "github.com/bradleyshawkins/rent/services/user/internal/identity"

	"github.com/matryer/is"
)

type mockUserCreatorService struct {
	RegisterUserUserRegistration       *identity2.User
	RegisterUserCredentials            *identity2.Credentials
	RegisterUserError                  error
	RegisterAccountUserID              identity2.UserID
	RegisterAccountAccountRegistration *identity2.Account
	RegisterAccountError               error
	AddUserToAccountAccountID          identity2.AccountID
	AddUserToAccountUserID             identity2.UserID
	AddUserToAccountRole               identity2.Role
	AddUserToAccountError              error
}

func (m *mockUserCreatorService) RegisterUser(u *identity2.User, c *identity2.Credentials) error {
	m.RegisterUserUserRegistration = u
	m.RegisterUserCredentials = c
	return m.RegisterUserError
}

func (m *mockUserCreatorService) RegisterAccount(a *identity2.Account) error {
	m.RegisterAccountAccountRegistration = a
	return m.RegisterAccountError
}

func (m *mockUserCreatorService) AddUserToAccount(aID identity2.AccountID, pID identity2.UserID, role identity2.Role) error {
	m.AddUserToAccountAccountID = aID
	m.AddUserToAccountUserID = pID
	m.AddUserToAccountRole = role
	return m.AddUserToAccountError
}

type mockUserCreator struct {
	mpcs *mockUserCreatorService
}

func (m *mockUserCreator) SignUp(suf *identity2.SignUpForm) error {
	return suf.SignUp(m.mpcs)
}

func TestRegisterUser(t *testing.T) {
	i := is.New(t)
	mpcs := &mockUserCreatorService{}
	mpc := &mockUserCreator{mpcs}
	registrar := identity2.NewSignUpManager(mpc)
	emailAddress, _ := mail.ParseAddress("email.address@test.com")
	firstName := "First"
	lastName := "Last"
	username := "username"
	password := "Password"

	user, err := registrar.SignUp(username, password, emailAddress, firstName, lastName)
	if err != nil {
		t.Fatal("Unexpected error occurred. Error:", err)
	}

	if user == nil {
		t.Fatal("signUpForm was nil.")
	}

	t.Log("user:", user)

	i.Equal(user.EmailAddress.Address, emailAddress.Address)
	i.Equal(user.FirstName, firstName)
	i.Equal(user.LastName, lastName)
	i.True(!user.ID.IsZero())

}

func TestRegisterUser_Fail(t *testing.T) {
	tests := []struct {
		name               string
		createUserError    error
		createAccountError error
	}{
		{
			name:               "Create user Fails",
			createUserError:    errors.New("unable to create user"),
			createAccountError: nil,
		},
		{
			name:               "Create account Fails",
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

			registrar := identity2.NewSignUpManager(mpc)

			emailAddress, _ := mail.ParseAddress("email.address@test.com")

			user, err := registrar.SignUp("username", "password", emailAddress, "First", "Last")
			if err == nil {
				t.Error("Received a nil error. Should have received an error")
			}
			if user != nil {
				t.Error("user was non-nil. user:", user)
			}
			t.Logf("error: %v", err)
		})
	}
}
