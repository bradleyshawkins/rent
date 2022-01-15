package identity

import (
	"fmt"
	"net/mail"

	"github.com/bradleyshawkins/rent"
)

type UserRegistration struct {
	ID           UserID
	EmailAddress *mail.Address
	Password     string
	FirstName    string
	LastName     string
	Status       UserStatus
}

type AccountRegistration struct {
	ID     AccountID
	Status AccountStatus
}

// RegistrationService contains all methods used to register a user.
// It will typically be implemented by a transaction that allows all calls to be chained together
type RegistrationService interface {
	RegisterUser(u *UserRegistration) error
	RegisterAccount(userID UserID, a *AccountRegistration) error
	AddUserToAccount(accountID AccountID, u UserID, role Role) error
}

// registrar is the interface that begins the registration process.
// Typically, it will be implemented by the database
type registrar interface {
	Register(registerFunc RegistrationFunc) error
}

// Registrar handles registering a user and creating an account for them
type Registrar struct {
	uc registrar
}

// NewRegistrar is a constructor for Registrar
func NewRegistrar(uc registrar) *Registrar {
	return &Registrar{uc: uc}
}

// RegistrationFunc is a closer around the steps used to register a user
type RegistrationFunc func(us RegistrationService) error

// Register registers a user and creates an account for them
func (u *Registrar) Register(emailAddress string, firstName, lastName string, password string) (*UserRegistration, *AccountRegistration, error) {
	addr, err := mail.ParseAddress(emailAddress)
	if err != nil {
		return nil, nil, err
	}

	user := &UserRegistration{
		ID:           NewUserID(),
		EmailAddress: addr,
		Password:     password,
		FirstName:    firstName,
		LastName:     lastName,
		Status:       UserActive,
	}

	account := &AccountRegistration{
		ID:     NewAccountID(),
		Status: AccountActive,
	}

	err = u.uc.Register(u.register(user, account))
	if err != nil {
		return nil, nil, err
	}

	return user, account, nil
}

// register is a closure that registers the user and an account
func (u *Registrar) register(user *UserRegistration, account *AccountRegistration) RegistrationFunc {
	return func(us RegistrationService) error {
		err := us.RegisterUser(user)
		if err != nil {
			return err
		}

		err = us.RegisterAccount(user.ID, account)
		if err != nil {
			return err
		}

		return nil
	}
}

func (u *Registrar) RegisterUserToAccount(accountID AccountID, role string, emailAddress string, firstName string, lastName string, password string) (*UserRegistration, error) {
	addr, err := mail.ParseAddress(emailAddress)
	if err != nil {
		return nil, err
	}

	r, ok := roleMap[role]
	if !ok {
		return nil, rent.NewError(fmt.Errorf("invalid role provided. Role %v", role), rent.WithInvalidPayload())
	}

	user := &UserRegistration{
		ID:           NewUserID(),
		EmailAddress: addr,
		Password:     password,
		FirstName:    firstName,
		LastName:     lastName,
		Status:       UserActive,
	}

	err = u.uc.Register(u.registerUserToAccount(accountID, user, r))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *Registrar) registerUserToAccount(accountID AccountID, user *UserRegistration, role Role) RegistrationFunc {
	return func(us RegistrationService) error {
		err := us.RegisterUser(user)
		if err != nil {
			return err
		}

		err = us.AddUserToAccount(accountID, user.ID, role)
		if err != nil {
			return err
		}
		return nil
	}
}
