package rent

import (
	uuid "github.com/satori/go.uuid"
)

type Renter struct {
	Person
	AccountID            uuid.UUID
	House                *Property
	ApprovedApplications []*Property
	Applications         []*Property
}
