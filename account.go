package rent

import uuid "github.com/satori/go.uuid"

type Account struct {
	ID     uuid.UUID
	Status AccountStatus
}

type AccountStatus int

const (
	AccountActive AccountStatus = iota + 1
	AccountDisabled
	AccountCanceled
)

func NewAccount() *Account {
	return &Account{
		ID:     uuid.NewV4(),
		Status: AccountActive,
	}
}

type Role int

const (
	RoleOwner Role = iota + 1
	RoleWriter
	RoleReader
)
