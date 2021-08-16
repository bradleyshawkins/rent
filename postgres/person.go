package postgres

import (
	"github.com/bradleyshawkins/rent/account"
	uuid "github.com/satori/go.uuid"
)

//
func (p *Postgres) RegisterPerson(person *account.Person) error {
	tx, err := p.db.Begin()
	if err != nil {
		return convertToError(err, "unable to begin person registration transaction")
	}

	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO person(id, email_address, password, status_id) VALUES ($1, $2, $3, $4)", person.ID, person.EmailAddress, person.Password, person.Status)
	if err != nil {
		return convertToError(err, "unable to insert person")
	}

	// TODO: Is having a uuid the best idea for a pk for this table?
	_, err = tx.Exec(`INSERT INTO person_details(id, person_id, first_name, last_name) VALUES ($1, $2, $3, $4)`, uuid.NewV4(), person.ID, person.FirstName, person.LastName)
	if err != nil {
		return convertToError(err, "unable to insert person details")
	}

	err = tx.Commit()
	if err != nil {
		return convertToError(err, "unable to commit new person registration")
	}
	return nil
}
