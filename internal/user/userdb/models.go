package userdb

import (
	"net/mail"

	"github.com/bradleyshawkins/rent/internal/user"
	uuid "github.com/satori/go.uuid"
)

type userDB struct {
	ID           uuid.UUID   `db:"id"`
	Status       user.Status `db:"status"`
	FirstName    string      `db:"first_name"`
	LastName     string      `db:"last_name"`
	EmailAddress string      `db:"email_address"`
	Role         user.Role   `db:"role"`
}

func toUserDB(u *user.User) *userDB {
	return &userDB{
		ID:           u.ID,
		Status:       u.Status,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		EmailAddress: u.EmailAddress.String(),
		Role:         u.Role,
	}
}

func toUser(u *userDB) *user.User {
	addr, _ := mail.ParseAddress(u.EmailAddress)
	return &user.User{
		ID:           u.ID,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		EmailAddress: addr,
		Status:       u.Status,
		Role:         u.Role,
	}
}
