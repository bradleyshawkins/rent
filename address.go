package rent

import uuid "github.com/satori/go.uuid"

type Address struct {
	ID      uuid.UUID `db:"id"`
	Street1 string    `db:"street_1"`
	Street2 string    `db:"street_2"`
	City    string    `db:"city"`
	State   string    `db:"state"`
	Zipcode string    `db:"zipcode"`
}
