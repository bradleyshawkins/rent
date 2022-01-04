package postgres

import (
	"github.com/bradleyshawkins/rent/identity"
)

func (t *transaction) RegisterAccount(pID identity.PersonID, a *identity.AccountRegistration) error {
	_, err := t.tx.Exec(`INSERT INTO account(id, status) VALUES ($1, $2)`, a.ID, a.Status)
	if err != nil {
		return toRentError(err)
	}

	err = t.AddPersonToAccount(a.ID, pID, identity.RoleOwner)
	if err != nil {
		return err
	}
	return nil
}

func (t *transaction) AddPersonToAccount(aID identity.AccountID, pID identity.PersonID, role identity.Role) error {
	_, err := t.tx.Exec(`INSERT INTO membership(account_id, person_id, role_id) VALUES ($1, $2, $3)`, aID, pID, role)
	if err != nil {
		return toRentError(err)
	}
	return nil
}

//func registerAccount(db dbConn, a *rent.Account) error {
//	_, err := db.Exec(`INSERT INTO account(id, status) VALUES ($1, $2)`, a.ID, a.Status)
//	if err != nil {
//		return toRentError(err)
//	}
//	return nil
//}
//
//func addToAccount(db dbConn, aID uuid.UUID, per *rent.Person) error {
//	_, err := db.Exec(`INSERT INTO membership(account_id, person_id, role_id) VALUES ($1, $2, $3)`, aID, per.ID, rent.RoleOwner)
//	if err != nil {
//		return toRentError(err)
//	}
//	return nil
//}
//
//func removeFromAccount(db dbConn, aID uuid.UUID, pID uuid.UUID) error {
//	_, err := db.Exec(`DELETE FROM membership WHERE account_id = $1 AND person_id = $2`, aID, pID)
//	if err != nil {
//		return toRentError(err)
//	}
//	return nil
//}
