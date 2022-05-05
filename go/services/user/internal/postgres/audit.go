package postgres

import (
	"context"

	"github.com/bradleyshawkins/rent/services/user/internal/identity"

	uuid "github.com/satori/go.uuid"

	"github.com/bradleyshawkins/rent/datastore"
)

type action string

const (
	actionCreate action = "CREATE"
	actionUpdate action = "UPDATE"
	actionDelete action = "DELETE"
)

type objectType string

const (
	objectTypeUser       objectType = "User"
	objectTypeAccount    objectType = "Account"
	objectTypeMembership objectType = "Membership"
)

const (
	addAuditLogQuery = `INSERT INTO audit(id, user_id, object_type, object_value, action, old_value, new_value) 
							VALUES ($1, $2, $3, $4, $5, $6, $7)`
)

func (t *transaction) addAuditLog(ctx context.Context, userID identity.UserID, objectType objectType, objectValue string,
	action action, oldValue string, newValue string) error {
	_, err := t.tx.ExecContext(ctx, addAuditLogQuery, uuid.NewV4(), userID, objectType, objectValue, action, oldValue, newValue)
	if err != nil {
		return datastore.ToBError(err)
	}
	return nil
}
