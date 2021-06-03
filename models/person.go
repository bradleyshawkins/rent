package models

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

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
	if p.EmailAddress == "" {
		return fmt.Errorf("email address is required for users")
	}

	if p.FirstName == "" {
		return fmt.Errorf("first name is required for users")
	}

	if p.LastName == "" {
		return fmt.Errorf("last name is required for users")
	}

	return nil
}
