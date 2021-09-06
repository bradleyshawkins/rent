package postgres

import (
	"database/sql"

	"github.com/bradleyshawkins/rent"

	uuid "github.com/satori/go.uuid"
)

//
func (p *Postgres) RegisterPerson(person *rent.Person) error {
	tx, err := p.db.Begin()
	if err != nil {
		return rent.NewError(err, rent.WithInternal(), rent.WithMessage("unable to begin transaction to insert person"))
	}

	defer func() {
		err = tx.Rollback()
		if err != nil {
			err = rent.NewError(err, rent.WithInternal(), rent.WithMessage("unable to rollback transaction for inserting person"))
		}
	}()

	if err := p.insertPerson(tx, person); err != nil {
		return err
	}

	if err := p.insertPersonDetails(tx, person); err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return rent.NewError(err, rent.WithInternal(), rent.WithMessage("unable to commit transaction for inserting person"))
	}
	return nil
}

func (p *Postgres) insertPerson(tx *sql.Tx, person *rent.Person) error {
	_, err := tx.Exec("INSERT INTO person(id, email_address, password, status_id) VALUES ($1, $2, $3, $4)", person.ID, person.EmailAddress, person.Password, person.Status)
	if err != nil {
		return toRentError(err)
	}

	return err
}

func (p *Postgres) insertPersonDetails(tx *sql.Tx, person *rent.Person) error {
	// TODO: Is having a uuid the best idea for a pk for this table?
	_, err := tx.Exec(`INSERT INTO person_details(id, person_id, first_name, last_name) VALUES ($1, $2, $3, $4)`, uuid.NewV4(), person.ID, person.FirstName, person.LastName)
	if err != nil {
		return toRentError(err)
	}

	return err
}
