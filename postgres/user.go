package postgres

import (
	"net/mail"

	"github.com/bradleyshawkins/rent"
	"github.com/bradleyshawkins/rent/identity"

	uuid "github.com/satori/go.uuid"
)

func (d *Database) LoadUser(uID identity.UserID) (*identity.User, error) {
	var firstName, lastName, emailAddress string
	var status identity.UserStatus
	err := d.db.QueryRow(`SELECT aud.first_name, aud.last_name, aud.email_address, au.status
							FROM app_user au
							INNER JOIN app_user_details aud ON au.app_user_details_id = aud.id
							WHERE au.id = $1`, uID.AsUUID()).Scan(&firstName, &lastName, &emailAddress, &status)
	if err != nil {
		return nil, toRentError(err)
	}

	addr, err := mail.ParseAddress(emailAddress)
	if err != nil {
		return nil, rent.NewError(err, rent.WithInternal(), rent.WithMessage("unable to parse email address on person load"))
	}

	return &identity.User{
		ID:           uID,
		EmailAddress: addr,
		FirstName:    firstName,
		LastName:     lastName,
		Status:       status,
	}, nil
}

// SignUp provides a transaction around the sign up process
func (d *Database) SignUp(suf *identity.SignUpForm) error {
	tx, err := d.begin()
	if err != nil {
		return err
	}

	defer func() {
		err = tx.rollback()
	}()

	err = suf.SignUp(tx)
	if err != nil {
		return err
	}

	err = tx.commit()
	if err != nil {
		return err
	}
	return nil
}

func (t *transaction) RegisterUser(user *identity.User, c *identity.Credentials) error {
	detailsID := uuid.NewV4()
	_, err := t.tx.Exec(`INSERT INTO app_user_details(id, first_name, last_name, email_address) VALUES ($1, $2, $3, $4)`, detailsID, user.FirstName, user.LastName, user.EmailAddress.Address)
	if err != nil {
		return toRentError(err)
	}

	credentialsID := uuid.NewV4()
	_, err = t.tx.Exec(`INSERT INTO app_user_credentials(id, username, password) VALUES ($1, $2, $3)`, credentialsID, c.Username, c.Password)
	if err != nil {
		return toRentError(err)
	}

	_, err = t.tx.Exec("INSERT INTO app_user(id, status, app_user_credentials_id, app_user_details_id) VALUES ($1, $2, $3, $4)", user.ID.AsUUID(), user.Status, credentialsID, detailsID)
	if err != nil {
		return toRentError(err)
	}

	return nil
}
