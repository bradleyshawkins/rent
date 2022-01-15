package identity

import (
	uuid "github.com/satori/go.uuid"
)

type UserID uuid.UUID

func NewUserID() UserID {
	return UserID(uuid.NewV4())
}

func AsUserID(id uuid.UUID) UserID {
	return UserID(id)
}

func (p UserID) IsZero() bool {
	return p.AsUUID() == uuid.Nil
}

func (p UserID) AsUUID() uuid.UUID {
	return uuid.UUID(p)
}

func (p UserID) String() string {
	return p.AsUUID().String()
}

type UserStatus string

const (
	UserDisabled UserStatus = "Disabled"
	UserActive   UserStatus = "Active"
)
