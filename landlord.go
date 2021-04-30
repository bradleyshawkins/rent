package rent

import uuid "github.com/satori/go.uuid"

type Landlord struct {
	ID     uuid.UUID `db:"id"`
	Person Person
}
