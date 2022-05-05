package postgres

import (
	"context"

	"github.com/bradleyshawkins/rent/services/user/internal/identity"
)

func (t *transaction) RegisterAccount(ctx context.Context, userID identity.UserID, a *identity.Account) error {
	_, err := t.tx.ExecContext(ctx, `INSERT INTO account(id, status) VALUES ($1, $2)`, a.ID.AsUUID(), a.Status)
	if err != nil {
		return toBError(err)
	}

	err = t.addAuditLog(ctx, userID, objectTypeAccount, a.ID.String(), actionCreate, "", a.ID.String())
	if err != nil {
		return err
	}

	return nil
}

func (t *transaction) AddUserToAccount(ctx context.Context, aID identity.AccountID, uID identity.UserID, role identity.Role) error {
	_, err := t.tx.ExecContext(ctx, `INSERT INTO membership(account_id, app_user_id, role) VALUES ($1, $2, $3)`,
		aID.AsUUID(), uID.AsUUID(), role)
	if err != nil {
		return toBError(err)
	}

	err = t.addAuditLog(ctx, uID, objectTypeMembership, "", actionCreate, "", uID.String())
	if err != nil {
		return err
	}
	return nil
}
