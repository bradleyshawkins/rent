package identity

import (
	"fmt"
	"net/mail"

	"github.com/bradleyshawkins/rent"
)

type PersonRegistration struct {
	ID           PersonID
	EmailAddress *mail.Address
	Password     string
	FirstName    string
	LastName     string
	Status       PersonStatus
}

type AccountRegistration struct {
	ID     AccountID
	Status AccountStatus
}

// RegistrationService contains all methods used to register a person.
// It will typically be implemented by a transaction that allows all calls to be chained together
type RegistrationService interface {
	RegisterPerson(u *PersonRegistration) error
	RegisterAccount(personID PersonID, a *AccountRegistration) error
	AddPersonToAccount(accountID AccountID, u PersonID, role Role) error
}

// registrar is the interface that begins the registration process.
// Typically, it will be implemented by the database
type registrar interface {
	Register(registerFunc RegistrationFunc) error
}

// Registrar handles registering a person and creating an account for them
type Registrar struct {
	uc registrar
}

// NewRegistrar is a constructor for Registrar
func NewRegistrar(uc registrar) *Registrar {
	return &Registrar{uc: uc}
}

// RegistrationFunc is a closer around the steps used to register a person
type RegistrationFunc func(us RegistrationService) error

// Register registers a person and creates an account for them
func (u *Registrar) Register(emailAddress string, firstName, lastName string, password string) (*PersonRegistration, *AccountRegistration, error) {
	addr, err := mail.ParseAddress(emailAddress)
	if err != nil {
		return nil, nil, err
	}

	person := &PersonRegistration{
		ID:           NewPersonID(),
		EmailAddress: addr,
		Password:     password,
		FirstName:    firstName,
		LastName:     lastName,
		Status:       PersonActive,
	}

	account := &AccountRegistration{
		ID:     NewAccountID(),
		Status: AccountActive,
	}

	err = u.uc.Register(u.register(person, account))
	if err != nil {
		return nil, nil, err
	}

	return person, account, nil
}

// register is a closure that registers the person and an account
func (u *Registrar) register(person *PersonRegistration, account *AccountRegistration) RegistrationFunc {
	return func(us RegistrationService) error {
		err := us.RegisterPerson(person)
		if err != nil {
			return err
		}

		err = us.RegisterAccount(person.ID, account)
		if err != nil {
			return err
		}

		return nil
	}
}

func (u *Registrar) RegisterPersonToAccount(accountID AccountID, role string, emailAddress string, firstName string, lastName string, password string) (*PersonRegistration, error) {
	addr, err := mail.ParseAddress(emailAddress)
	if err != nil {
		return nil, err
	}

	r, ok := roleMap[role]
	if !ok {
		return nil, rent.NewError(fmt.Errorf("invalid role provided. Role %v", role), rent.WithInvalidPayload())
	}

	person := &PersonRegistration{
		ID:           NewPersonID(),
		EmailAddress: addr,
		Password:     password,
		FirstName:    firstName,
		LastName:     lastName,
		Status:       PersonActive,
	}

	err = u.uc.Register(u.registerPersonToAccount(accountID, person, r))
	if err != nil {
		return nil, err
	}

	return person, nil
}

func (u *Registrar) registerPersonToAccount(accountID AccountID, person *PersonRegistration, role Role) RegistrationFunc {
	return func(us RegistrationService) error {
		err := us.RegisterPerson(person)
		if err != nil {
			return err
		}

		err = us.AddPersonToAccount(accountID, person.ID, role)
		if err != nil {
			return err
		}
		return nil
	}
}
