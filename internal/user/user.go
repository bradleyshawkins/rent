package user

import (
	"net/mail"

	uuid "github.com/satori/go.uuid"
)

type Role string

const (
	RoleAdmin Role = "Admin"
	RoleUser  Role = "User"
)

type Status string

const (
	Active   Status = "Active"
	Inactive Status = "Inactive"
)

type User struct {
	ID           uuid.UUID
	Account      *Account
	FirstName    string
	LastName     string
	EmailAddress *mail.Address
	Status       Status
	Role         Role
}

type NewUser struct {
	FirstName    string
	LastName     string
	EmailAddress string
}

type UpdateUser struct {
	ID           uuid.UUID
	FistName     string
	LastName     string
	EmailAddress string
	Status       string
}
