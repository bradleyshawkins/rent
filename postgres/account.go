package postgres

import (
	"github.com/bradleyshawkins/rent"
	uuid "github.com/satori/go.uuid"
)

func registerAccount(db dbConn, a *rent.Account) error {
	_, err := db.Exec(`INSERT INTO account(id, status) VALUES ($1, $2)`, a.ID, a.Status)
	if err != nil {
		return toRentError(err)
	}
	return nil
}

func addToAccount(db dbConn, aID uuid.UUID, per *rent.Person) error {
	_, err := db.Exec(`INSERT INTO membership(account_id, person_id, role_id) VALUES ($1, $2, $3)`, aID, per.ID, rent.RoleOwner)
	if err != nil {
		return toRentError(err)
	}
	return nil
}

func removeFromAccount(db dbConn, aID uuid.UUID, pID uuid.UUID) error {
	_, err := db.Exec(`DELETE FROM membership WHERE account_id = $1 AND person_id = $2`, aID, pID)
	if err != nil {
		return toRentError(err)
	}
	return nil
}
