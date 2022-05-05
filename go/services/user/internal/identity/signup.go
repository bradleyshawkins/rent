package identity

import (
	"context"
	"net/mail"

	"github.com/bradleyshawkins/rent/berror"

	"golang.org/x/crypto/bcrypt"
)

// signUpper is the interface that begins the signup process.
type signUpper interface {
	SignUp(ctx context.Context, s *SignUpForm) error
}

// SignUpForm contains all necessary data needed to sign up
type SignUpForm struct {
	Credentials *Credentials
	User        *User
	Account     *Account
}

// SignUpManager handles initiating the signup process
type SignUpManager struct {
	su signUpper
}

// NewSignUpManager is a constructor for SignUpManager
func NewSignUpManager(uc signUpper) *SignUpManager {
	return &SignUpManager{su: uc}
}

// SignUp creates the types needed for signing off and kicks off the signing up steps
func (u *SignUpManager) SignUp(ctx context.Context, username string, password string, emailAddress *mail.Address, firstName, lastName string) (*User, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), 6)
	if err != nil {
		return nil, berror.WrapInternal(err, "unable to encrypt password")
	}

	form := &SignUpForm{
		Credentials: &Credentials{
			Username: username,
			Password: string(pw),
		},
		User: &User{
			ID:           NewUserID(),
			EmailAddress: emailAddress,
			FirstName:    firstName,
			LastName:     lastName,
			Status:       UserActive,
		},
		Account: &Account{
			ID:     NewAccountID(),
			Status: AccountActive,
		},
	}

	err = u.su.SignUp(ctx, form)
	if err != nil {
		return nil, err
	}

	return form.User, nil
}
