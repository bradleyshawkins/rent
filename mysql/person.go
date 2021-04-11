package mysql

import (
	"database/sql"
	"fmt"

	uuid "github.com/satori/go.uuid"

	"github.com/jmoiron/sqlx"

	"github.com/bradleyshawkins/rent"
)

func (m *MySQL) RegisterPerson(p *rent.Person) error {
	tx, err := m.db.Beginx()
	if err != nil {
		return fmt.Errorf("unable to initialize transaction. Error: %v", err)
	}

	err = m.insertPerson(tx, p)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("unable to insert person. Error: %v", err)
	}

	return tx.Commit()
}

func (m *MySQL) insertPerson(tx *sqlx.Tx, p *rent.Person) error {
	_, err := tx.NamedExec("INSERT INTO person(id, first_name, last_name, email_address) VALUES (:id, :first_name, :last_name, :email_address)", p)
	if err != nil {
		return fmt.Errorf("unable to insert person into database. Error: %v", err)
	}
	return nil
}

func (m *MySQL) UpdatePerson(t *rent.Person) error {
	return nil
}

func (m *MySQL) GetPerson(id uuid.UUID) (*rent.Person, error) {
	var person rent.Person
	err := m.db.Get(&person, `SELECT id, first_name, last_name, email_address FROM person WHERE id=?`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("unable to get person. Error: %v", err)
	}
	return &person, nil
}

func (m *MySQL) DeletePerson(id uuid.UUID) error {
	return nil
}
