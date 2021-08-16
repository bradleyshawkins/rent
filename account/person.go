package account

import (
	"regexp"

	"github.com/bradleyshawkins/rent/types"
	uuid "github.com/satori/go.uuid"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Person struct {
	ID           uuid.UUID
	EmailAddress string
	Password     string
	FirstName    string
	LastName     string
	Status       PersonStatus

	as accountService
}

type PersonStatus int

const (
	PersonActive PersonStatus = iota + 1
	PersonInactive
	PersonDisabled
)

func NewPerson(as accountService, emailAddress, password, firstName, lastName string) (*Person, error) {
	p := &Person{
		ID:           uuid.NewV4(),
		EmailAddress: emailAddress,
		Password:     password,
		FirstName:    firstName,
		LastName:     lastName,
		Status:       PersonActive,

		as: as,
	}
	return p, p.validate()
}

func (p Person) validate() error {

	if p.as == nil {
		return types.NewSetupError("accountService", types.SetupNotSet)
	}

	if p.ID == (uuid.UUID{}) {
		return types.NewFieldValidationError("id", types.Missing)
	}

	if p.Password == "" {
		return types.NewFieldValidationError("password", types.Missing)
	}

	if p.FirstName == "" {
		return types.NewFieldValidationError("firstName", types.Missing)
	}

	if p.LastName == "" {
		return types.NewFieldValidationError("lastName", types.Missing)
	}

	if err := validateEmailAddress(p.EmailAddress); err != nil {
		return err
	}

	return nil
}

func validateEmailAddress(email string) error {
	if email == "" {
		return types.NewFieldValidationError("emailAddress", types.Missing)
	}

	if !emailRegex.MatchString(email) {
		return types.NewFieldValidationError("emailAddress", types.Invalid)
	}

	return nil
}

func (p *Person) Register() error {
	if err := p.validate(); err != nil {
		return err
	}

	err := p.as.RegisterPerson(p)
	if err != nil {
		return err
	}
	return nil
}
