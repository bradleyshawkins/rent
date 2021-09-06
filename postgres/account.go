package postgres

import (
	"database/sql"

	"github.com/bradleyshawkins/rent"
	uuid "github.com/satori/go.uuid"
)

func (p *Postgres) LoadAccount(accountID uuid.UUID) (*rent.Account, error) {
	var status rent.AccountStatus
	err := p.db.QueryRow(`SELECT status_id FROM account WHERE id = $1`, accountID).Scan(&status)
	if err != nil {
		return nil, convertToError(err, "unable to get account")
	}
	return rent.NewAccountWithID(accountID, status)
}

func (p *Postgres) registerAccount(tx *sql.Tx) (uuid.UUID, error) {
	id := uuid.NewV4()
	_, err := tx.Exec(`INSERT INTO account(id, status_id) VALUES ($1, $2)`, id, rent.AccountActive)
	if err != nil {
		return uuid.UUID{}, convertToError(err, "unable to insert account")
	}
	return id, nil
}

func (p *Postgres) AddToAccount(aID uuid.UUID, pe *rent.Person) error {
	_, err := p.db.Exec(`INSERT INTO membership(person_id, account_id, permission) VALUES ($1, $2, $3)`, pe.ID, aID, 1)
	if err != nil {
		return convertToError(err, "unable to add person to account")
	}
	return nil
}

func (p *Postgres) addToAccount(tx *sql.Tx, pe *rent.Person) error {
	_, err := tx.Exec(`INSERT INTO membership(person_id, account_id, permission) VALUES ($1, $2, $3)`, pe.ID, uuid.NewV4(), 1)
	if err != nil {
		return convertToError(err, "unable to add person to account")
	}
	return nil
}
