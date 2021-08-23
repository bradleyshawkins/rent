package postgres

import (
	"github.com/bradleyshawkins/rent/account"
	uuid "github.com/satori/go.uuid"
)

func (p *Postgres) LoadAccount(accountID uuid.UUID) (*account.Account, error) {
	var status account.AccountStatus
	err := p.db.QueryRow(`SELECT status_id FROM account WHERE id = $1`, accountID).Scan(&status)
	if err != nil {
		return nil, convertToError(err, "unable to get account")
	}
	return account.NewAccountWithID(p, accountID, status)
}

func (p *Postgres) RegisterAccount(a *account.Account) error {
	_, err := p.db.Exec(`INSERT INTO account(id, status_id) VALUES ($1, $2)`, a.ID, a.Status)
	if err != nil {
		return convertToError(err, "unable to insert account")
	}
	return nil
}

func (p *Postgres) AddToAccount(aID uuid.UUID, pe *account.Person) error {
	_, err := p.db.Exec(`INSERT INTO membership(person_id, account_id, permission) VALUES ($1, $2, $3)`, pe.ID, aID, 1)
	if err != nil {
		return convertToError(err, "unable to add person to account")
	}
	return nil
}
