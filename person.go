package rent

import (
	"errors"
	"regexp"

	uuid "github.com/satori/go.uuid"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type PersonStore interface {
	RegisterPerson(a *Account, p *Person) error
	LoadPerson(id uuid.UUID) (*Person, error)
	CancelPerson(accountID, personID uuid.UUID) error
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
	PersonDisabled PersonStatus = iota + 1
	PersonActive
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

func NewExistingPerson(id uuid.UUID, emailAddress, password, firstName, lastName string, status PersonStatus) (*Person, error) {
	p := &Person{
		ID:           id,
		EmailAddress: emailAddress,
		Password:     password,
		FirstName:    firstName,
		LastName:     lastName,
		Status:       status,
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
		return NewError(errors.New("missing password"), WithInvalidFields(InvalidField{
			Field:  "password",
			Reason: ReasonMissing,
		}))
	}

	if p.FirstName == "" {
		return NewError(errors.New("missing firstName"), WithInvalidFields(InvalidField{
			Field:  "firstName",
			Reason: ReasonMissing,
		}))
	}

	if p.LastName == "" {
		return NewError(errors.New("missing lastName"), WithInvalidFields(InvalidField{
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
		return NewError(errors.New("missing email"), WithInvalidFields(InvalidField{
			Field:  "emailAddress",
			Reason: ReasonMissing,
		}))
	}

	if !emailRegex.MatchString(email) {
		return NewError(errors.New("invalid email"), WithInvalidFields(InvalidField{
			Field:  "emailAddress",
			Reason: ReasonInvalid,
		}))
	}

	return nil
}

func (p *Person) IsActive() error {
	if p.Status == PersonDisabled {
		return NewError(errors.New("person is not active"), WithEntityDisabled())
	}
	return nil
}

func (p *Person) Disable() {
	p.Status = PersonDisabled
}
