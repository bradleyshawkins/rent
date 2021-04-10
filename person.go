package rent

import uuid "github.com/satori/go.uuid"

type Person struct {
	ID            uuid.UUID `db:"id"`
	FirstName     string    `db:"first_name"`
	MiddleInitial string    `db:"middle_initial"`
	LastName      string    `db:"last_name"`
	EmailAddress  string    `db:"email_address"`
	Address       Address
}
