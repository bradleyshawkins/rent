package postgres

import (
	"github.com/bradleyshawkins/rent/identity"

	uuid "github.com/satori/go.uuid"
)

// Register provides a transaction around registering accounts and users
func (d *Database) Register(registrationFunc identity.RegistrationFunc) error {
	tx, err := d.begin()
	if err != nil {
		return err
	}

	defer tx.rollback()

	err = registrationFunc(tx)
	if err != nil {
		return err
	}

	err = tx.commit()
	if err != nil {
		return err
	}
	return nil
}

func (t *transaction) RegisterUser(user *identity.UserRegistration) error {
	detailsID := uuid.NewV4()
	_, err := t.tx.Exec(`INSERT INTO app_user_details(id, first_name, last_name) VALUES ($1, $2, $3)`, detailsID, user.FirstName, user.LastName)
	if err != nil {
		return toRentError(err)
	}

	_, err = t.tx.Exec("INSERT INTO app_user(id, email_address, password, status, app_user_details_id) VALUES ($1, $2, $3, $4, $5)", user.ID.AsUUID(), user.EmailAddress.Address, user.Password, user.Status, detailsID)
	if err != nil {
		return toRentError(err)
	}

	return err
}

//
//// RegisterUser inserts the person into the database, creates an account and associates the person with the account
//func (p *Postgres) RegisterUser(a *rent.Account, person *rent.Person) error {
//	if err := person.Validate(); err != nil {
//		return err
//	}
//
//	tx, err := p.db.Begin()
//	if err != nil {
//		return rent.NewError(err, rent.WithInternal(), rent.WithMessage("unable to begin transaction to insert person"))
//	}
//
//	defer tx.Rollback()
//
//	if err := insertPerson(tx, person); err != nil {
//		return err
//	}
//
//	err = insertPersonDetails(tx, person)
//	if err != nil {
//		return err
//	}
//
//	err = registerAccount(tx, a)
//	if err != nil {
//		return err
//	}
//
//	err = addToAccount(tx, a.ID, person)
//	if err != nil {
//		return err
//	}
//
//	err = tx.Commit()
//	if err != nil {
//		return rent.NewError(err, rent.WithInternal(), rent.WithMessage("unable to commit transaction for inserting person"))
//	}
//	return nil
//}
//
//func (p *Postgres) LoadPerson(id uuid.UUID) (*rent.Person, error) {
//	person, err := getPersonByID(p.db, id)
//	if err != nil {
//		return nil, err
//	}
//
//	return person, person.IsActive()
//}
//
//func (p *Postgres) CancelPerson(accountID, personID uuid.UUID) error {
//	tx, err := p.db.Begin()
//	if err != nil {
//		return rent.NewError(err, rent.WithInternal(), rent.WithMessage("unable to begin transaction to insert person"))
//	}
//
//	defer tx.Rollback()
//
//	person, err := getPersonByID(tx, personID)
//	if err != nil {
//		return err
//	}
//
//	person.Disable()
//
//	err = updatePersonStatus(tx, personID, person.Status)
//	if err != nil {
//		return err
//	}
//
//	err = removeFromAccount(tx, accountID, personID)
//	if err != nil {
//		return err
//	}
//
//	err = tx.Commit()
//	if err != nil {
//		err = rent.NewError(err, rent.WithInternal(), rent.WithMessage("unable to commit transaction for canceling person"))
//	}
//
//	return nil
//}
//
//func insertPerson(conn dbConn, person *rent.Person) error {
//	_, err := conn.Exec("INSERT INTO person(id, email_address, password, status_id) VALUES ($1, $2, $3, $4)", person.ID, person.EmailAddress, person.Password, person.Status)
//	if err != nil {
//		return toRentError(err)
//	}
//
//	return err
//}
//
//func insertPersonDetails(conn dbConn, person *rent.Person) error {
//	detailsID := uuid.NewV4()
//	_, err := conn.Exec(`INSERT INTO person_details(id, person_id, first_name, last_name) VALUES ($1, $2, $3, $4)`, detailsID, person.ID, person.FirstName, person.LastName)
//	if err != nil {
//		return toRentError(err)
//	}
//
//	return err
//}
//
//func getPersonByID(conn dbConn, id uuid.UUID) (*rent.Person, error) {
//	var emailAddress, password, firstName, lastName string
//	var statusID rent.UserStatus
//	err := conn.QueryRow(`SELECT p.email_address, p.password, p.status_id, pd.first_name, pd.last_name
//										FROM person p
//										INNER JOIN person_details pd ON p.id = pd.person_id
//										WHERE p.id = $1`, id).Scan(&emailAddress, &password, &statusID,
//		&firstName, &lastName)
//	if err != nil {
//		return nil, toRentError(err)
//	}
//
//	return rent.NewExistingPerson(id, emailAddress, password, firstName, lastName, statusID)
//}
//
//func updatePersonStatus(conn dbConn, id uuid.UUID, status rent.UserStatus) error {
//	_, err := conn.Exec(`UPDATE person SET status_id = $1 WHERE id = $2`, status, id)
//	if err != nil {
//		return toRentError(err)
//	}
//
//	return nil
//}
