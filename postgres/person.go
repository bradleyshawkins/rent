package postgres

import (
	"database/sql"
	"fmt"

	uuid "github.com/satori/go.uuid"

	"github.com/bradleyshawkins/rent"
)

func (m *Postgres) CreatePerson(p *rent.Person) error {
	_, err := m.db.NamedExec("INSERT INTO user(id, first_name, last_name, email_address) VALUES (:id, :first_name, :last_name, :email_address)", p)
	if err != nil {
		return convertToError(err, "unable to insert person into database")
	}
	return nil
}

func (m *Postgres) UpdatePerson(t *rent.Person) error {
	_, err := m.db.NamedExec(`UPDATE user SET id=:id, first_name=:first_name, 
                  					last_name=:last_name, email_address=:email_address WHERE id = :id`, t)
	if err != nil {
		return convertToError(err, "unable to update person")
	}
	return nil
}

func (m *Postgres) GetPerson(id uuid.UUID) (*rent.Person, error) {
	var user rent.Person
	err := m.db.Get(&user, `SELECT id, first_name, last_name, email_address FROM user WHERE id=$1`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("unable to get person. Error: %w", err)
	}
	return &user, nil
}

func (m *Postgres) DeletePerson(id uuid.UUID) error {
	_, err := m.db.Exec(`DELETE FROM user WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("unable to delete person from database. Error: %w", err)
	}
	return nil
}
