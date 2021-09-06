package rent

import (
	"errors"
	"regexp"

	uuid "github.com/satori/go.uuid"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type PersonService interface {
	RegisterPerson(p *Person) error
	//LoadPerson(id uuid.UUID) (*Person, error)
}

type Person struct {
	ID           uuid.UUID
	EmailAddress string
	Password     string
	FirstName    string
	LastName     string
	Status       PersonStatus
}

type PersonStatus int

const (
	PersonActive PersonStatus = iota + 1
	PersonInactive
	PersonDisabled
)

func NewPerson(emailAddress, password, firstName, lastName string) (*Person, error) {
	p := &Person{
		ID:           uuid.NewV4(),
		EmailAddress: emailAddress,
		Password:     password,
		FirstName:    firstName,
		LastName:     lastName,
		Status:       PersonActive,
	}
	return p, p.Validate()
}

func (p Person) Validate() error {
	if p.ID == (uuid.UUID{}) {
		return NewError(errors.New("missing id"), WithInvalidFields(InvalidField{
			Field:  "id",
			Reason: ReasonMissing,
		}))
	}

	if p.Password == "" {
		return NewError(errors.New("missing id"), WithInvalidFields(InvalidField{
			Field:  "password",
			Reason: ReasonMissing,
		}))
	}

	if p.FirstName == "" {
		return NewError(errors.New("missing id"), WithInvalidFields(InvalidField{
			Field:  "firstName",
			Reason: ReasonMissing,
		}))
	}

	if p.LastName == "" {
		return NewError(errors.New("missing id"), WithInvalidFields(InvalidField{
			Field:  "lastName",
			Reason: ReasonMissing,
		}))
	}

	if err := validateEmailAddress(p.EmailAddress); err != nil {
		return err
	}

	return nil
}

func validateEmailAddress(email string) error {
	if email == "" {
		return NewError(errors.New("missing id"), WithInvalidFields(InvalidField{
			Field:  "emailAddress",
			Reason: ReasonMissing,
		}))
	}

	if !emailRegex.MatchString(email) {
		return NewError(errors.New("missing id"), WithInvalidFields(InvalidField{
			Field:  "emailAddress",
			Reason: ReasonInvalid,
		}))
	}

	return nil
}
