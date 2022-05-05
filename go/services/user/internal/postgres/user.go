package postgres

import (
	"context"

	"github.com/bradleyshawkins/rent/datastore"

	"github.com/bradleyshawkins/rent/services/user/internal/identity"
)

const (
	insertUserQuery = `INSERT INTO app_user(id, status, first_name, last_name, email_address, username, password) 
							VALUES ($1, $2, $3, $4, $5, $6, $7)`
)

// SignUp provides a transaction around the sign up process
func (d *Database) SignUp(ctx context.Context, suf *identity.SignUpForm) error {
	tx, err := d.begin()
	if err != nil {
		return err
	}

	defer func() {
		err = tx.rollback()
	}()

	err = tx.RegisterUser(ctx, suf.User, suf.Credentials)
	if err != nil {
		return err
	}

	err = tx.RegisterAccount(ctx, suf.User.ID, suf.Account)
	if err != nil {
		return err
	}

	err = tx.AddUserToAccount(ctx, suf.Account.ID, suf.User.ID, identity.RoleOwner)
	if err != nil {
		return err
	}

	err = tx.commit()
	if err != nil {
		return err
	}
	return nil
}

func (t *transaction) RegisterUser(ctx context.Context, user *identity.User, c *identity.Credentials) error {
	_, err := t.tx.ExecContext(ctx, insertUserQuery, user.ID, user.Status, user.FirstName,
		user.LastName, user.EmailAddress, c.Username, c.Password)
	if err != nil {
		return datastore.ToBError(err)
	}

	err = t.addAuditLog(ctx, user.ID, objectTypeUser, user.ID.String(), actionCreate, "", user.ID.String())
	if err != nil {
		return err
	}

	return nil
}
