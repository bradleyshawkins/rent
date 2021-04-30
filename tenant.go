package rent

import uuid "github.com/satori/go.uuid"

type Tenant struct {
	ID     uuid.UUID `db:"id"`
	Person Person
}
