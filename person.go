package rent

import (
	"regexp"

	uuid "github.com/satori/go.uuid"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Person struct {
	ID           uuid.UUID
	Username     string
	Password     string
	FirstName    string
	LastName     string
	EmailAddress string
	PhoneNumber  string
}

func NewPerson(username, password, firstName, lastName, emailAddress, phoneNumber string) (*Person, error) {
	p := &Person{
		ID:           uuid.NewV4(),
		Username:     username,
		Password:     password,
		FirstName:    firstName,
		LastName:     lastName,
		PhoneNumber:  phoneNumber,
		EmailAddress: emailAddress,
	}
	return p, p.Validate()
}

func NewPersonWithID(id uuid.UUID, username, password, firstName, lastName, emailAddress, phoneNumber string) (*Person, error) {
	p := &Person{
		ID:           id,
		Username:     username,
		Password:     password,
		FirstName:    firstName,
		LastName:     lastName,
		PhoneNumber:  phoneNumber,
		EmailAddress: emailAddress,
	}
	return p, p.Validate()
}
func (p Person) Validate() error {

	if p.ID == (uuid.UUID{}) {
		return NewValidationError("id", Missing)
	}

	if p.Username == "" {
		return NewValidationError("username", Missing)
	}

	if p.Password == "" {
		return NewValidationError("password", Missing)
	}

	if p.FirstName == "" {
		return NewValidationError("firstName", Missing)
	}

	if p.LastName == "" {
		return NewValidationError("lastName", Missing)
	}

	if err := ValidateEmailAddress(p.EmailAddress); err != nil {
		return err
	}

	return nil
}

func ValidateEmailAddress(email string) error {
	if email == "" {
		return NewValidationError("emailAddress", Missing)
	}

	if !emailRegex.MatchString(email) {
		return NewValidationError("emailAddress", Invalid)
	}

	return nil
}
