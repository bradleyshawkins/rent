package postgres

import (
	"github.com/bradleyshawkins/rent"
	uuid "github.com/satori/go.uuid"
)

func registerAccount(db dbConn) (uuid.UUID, error) {
	a := rent.NewAccount()
	_, err := db.Exec(`INSERT INTO account(id, status) VALUES ($1, $2)`, a.ID, a.Status)
	if err != nil {
		return uuid.UUID{}, toRentError(err)
	}
	return a.ID, nil
}

func addToAccount(db dbConn, aID uuid.UUID, per *rent.Person) error {
	_, err := db.Exec(`INSERT INTO membership(account_id, person_id, role_id) VALUES ($1, $2, $3)`, aID, per.ID, rent.RoleOwner)
	if err != nil {
		return toRentError(err)
	}
	return nil
}
