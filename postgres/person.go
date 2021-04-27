package postgres

import (
	"database/sql"
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"

	"github.com/jmoiron/sqlx"

	"github.com/bradleyshawkins/rent"
)

func (m *Postgres) RegisterPerson(p *rent.Person) error {
	tx, err := m.db.Beginx()
	if err != nil {
		return fmt.Errorf("unable to initialize transaction. Error: %v", err)
	}

	err = m.insertPerson(tx, p)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("unable to insert person. Error: %w", err)
	}

	return tx.Commit()
}

func (m *Postgres) insertPerson(tx *sqlx.Tx, p *rent.Person) error {
	log.Printf("Inserting person: %+v\n", p)
	_, err := tx.NamedExec("INSERT INTO person(id, first_name, last_name, email_address) VALUES (:id, :first_name, :last_name, :email_address)", p)
	if err != nil {
		return convertToError(err, "unable to insert person into database")
	}
	return nil
}

func (m *Postgres) UpdatePerson(t *rent.Person) error {
	_, err := m.db.NamedExec(`UPDATE person SET id=:id, first_name=:first_name, 
                  					last_name=:last_name, email_address=:email_address WHERE id = :id`, t)
	if err != nil {
		return convertToError(err, "unable to update person")
	}
	return nil
}

func (m *Postgres) GetPerson(id uuid.UUID) (*rent.Person, error) {
	var person rent.Person
	err := m.db.Get(&person, `SELECT id, first_name, last_name, email_address FROM person WHERE id=?`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("unable to get person. Error: %w", err)
	}
	return &person, nil
}

func (m *Postgres) DeletePerson(id uuid.UUID) error {
	_, err := m.db.Exec(`DELETE FROM person WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("unable to delete person from database. Error: %w", err)
	}
	return nil
}
