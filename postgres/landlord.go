package postgres

import (
	"github.com/bradleyshawkins/rent"
	uuid "github.com/satori/go.uuid"
)

func (p *Postgres) RegisterLandlord(landlord *rent.Landlord) error {
	tx, err := p.db.Beginx()
	if err != nil {
		return convertToError(err, "unable to start transaction")
	}

	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO person(id, username, password, first_name, last_name, email_address) VALUES ($1, $2, $3, $4, $5, $6)", landlord.ID, landlord.Username, landlord.Password, landlord.FirstName, landlord.LastName, landlord.EmailAddress)
	if err != nil {
		return convertToError(err, "unable to insert person")
	}

	_, err = tx.Exec("INSERT INTO landlord(id, person_id) VALUES ($1, $2)", landlord.LandlordID, landlord.ID)

	err = tx.Commit()
	if err != nil {
		return convertToError(err, "unable to commit landlord registration")
	}

	return nil
}

func (p *Postgres) CancelLandlord(landlordID uuid.UUID) error {
	return nil
}

func (p *Postgres) AddProperty(property rent.Property) error {
	return nil
}
