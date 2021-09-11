package postgres

import (
	"github.com/bradleyshawkins/rent"

	uuid "github.com/satori/go.uuid"
)

// RegisterPerson inserts the person into the database, creates an account and associates the person with the account
func (p *Postgres) RegisterPerson(a *rent.Account, person *rent.Person) error {
	if err := person.Validate(); err != nil {
		return err
	}

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

	if err := insertPerson(tx, person); err != nil {
		return err
	}

	err = insertPersonDetails(tx, person)
	if err != nil {
		return err
	}

	err = registerAccount(tx, a)
	if err != nil {
		return err
	}

	err = addToAccount(tx, a.ID, person)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return rent.NewError(err, rent.WithInternal(), rent.WithMessage("unable to commit transaction for inserting person"))
	}
	return nil
}

func (p *Postgres) LoadPerson(id uuid.UUID) (*rent.Person, error) {
	return getPersonByID(p.db, id)
}

func insertPerson(conn dbConn, person *rent.Person) error {
	_, err := conn.Exec("INSERT INTO person(id, email_address, password, status_id) VALUES ($1, $2, $3, $4)", person.ID, person.EmailAddress, person.Password, person.Status)
	if err != nil {
		return toRentError(err)
	}

	return err
}

func insertPersonDetails(conn dbConn, person *rent.Person) error {
	detailsID := uuid.NewV4()
	_, err := conn.Exec(`INSERT INTO person_details(id, person_id, first_name, last_name) VALUES ($1, $2, $3, $4)`, detailsID, person.ID, person.FirstName, person.LastName)
	if err != nil {
		return toRentError(err)
	}

	return err
}

func getPersonByID(conn dbConn, id uuid.UUID) (*rent.Person, error) {
	var emailAddress, password, firstName, lastName string
	var statusID rent.PersonStatus
	err := conn.QueryRow(`SELECT p.email_address, p.password, p.status_id, pd.first_name, pd.last_name
										FROM person p
										INNER JOIN person_details pd ON p.id = pd.person_id
										WHERE p.id = $1`, id).Scan(&emailAddress, &password, &statusID,
		&firstName, &lastName)
	if err != nil {
		return nil, toRentError(err)
	}

	return rent.NewExistingPerson(id, emailAddress, password, firstName, lastName, statusID)
}
