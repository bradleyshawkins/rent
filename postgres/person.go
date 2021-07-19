package postgres

import (
	"github.com/bradleyshawkins/rent"
	uuid "github.com/satori/go.uuid"
)

func (p *Postgres) RegisterPerson(person *rent.Person) error {

	_, err := p.db.Exec("INSERT INTO person(id, username, password, first_name, last_name, email_address) VALUES ($1, $2, $3, $4, $5, $6)", person.ID, person.Username, person.Password, person.FirstName, person.LastName, person.EmailAddress)
	if err != nil {
		return convertToError(err, "unable to insert person")
	}

	return nil
}

func (p *Postgres) CancelPerson(personID uuid.UUID) error {
	_, err := p.db.Exec(`UPDATE person SET is_active = false WHERE id = $1`, personID)
	if err != nil {
		return convertToError(err, "unable to cancel landlord")
	}
	return nil
}

func (p *Postgres) GetPerson(personID uuid.UUID) (*rent.Person, error) {
	var username, firstName, lastName, emailAddress string
	var id uuid.UUID
	err := p.db.QueryRow(`SELECT id, username, first_name, last_name, email_address FROM person WHERE id = $1`, personID).Scan(&id, &username, &firstName, &lastName, &emailAddress)
	if err != nil {
		return nil, convertToError(err, "unable to select person")
	}
	return &rent.Person{
		ID:           id,
		Username:     username,
		Password:     "",
		FirstName:    firstName,
		LastName:     lastName,
		EmailAddress: emailAddress,
		PhoneNumber:  "",
	}, nil
}
