package identity

import uuid "github.com/satori/go.uuid"

type AccountID uuid.UUID

func NewAccountID() AccountID {
	return AsAccountID(uuid.NewV4())
}

func AsAccountID(id uuid.UUID) AccountID {
	return AccountID(id)
}

func (a AccountID) AsUUID() uuid.UUID {
	return uuid.UUID(a)
}

func (a AccountID) IsZero() bool {
	return a.AsUUID() == uuid.Nil
}

func (a AccountID) String() string {
	return a.AsUUID().String()
}

type AccountStatus string

const (
	AccountDisabled AccountStatus = "Disabled"
	AccountActive   AccountStatus = "Active"
	AccountCanceled AccountStatus = "Canceled"
)

type Role string

const (
	RoleOwner  Role = "Owner"
	RoleWriter Role = "Writer"
	RoleReader Role = "Reader"
)

var roleMap = map[string]Role{
	"Owner":  RoleOwner,
	"Writer": RoleWriter,
	"Reader": RoleReader,
}

type Account struct {
	ID     AccountID
	Status AccountStatus
}
