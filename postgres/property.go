package postgres

import (
	"github.com/bradleyshawkins/rent"
	uuid "github.com/satori/go.uuid"
)

func (p *Postgres) RegisterProperty(accountID uuid.UUID, prop *rent.Property) error {
	tx, err := p.db.Begin()
	if err != nil {
		return toRentError(err)
	}

	defer func() {
		err = tx.Rollback()
		if err != nil {
			err = rent.NewError(err, rent.WithInternal(), rent.WithMessage("unable to rollback transaction for inserting property"))
		}
	}()

	err = createProperty(tx, accountID, prop)
	if err != nil {
		return err
	}

	addrID, err := createAddress(tx, prop.Address)
	if err != nil {
		return err
	}

	err = addPropertyAddress(tx, prop.ID, addrID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return toRentError(err)
	}
	return nil
}

func createProperty(db dbConn, accountID uuid.UUID, prop *rent.Property) error {
	_, err := db.Exec(`INSERT INTO property(id, account_id, name, property_status_id) VALUES ($1, $2, $3, $4)`, prop.ID, accountID, prop.Name, prop.Status)
	if err != nil {
		return toRentError(err)
	}
	return nil
}

func createAddress(db dbConn, address *rent.Address) (uuid.UUID, error) {
	addressID := uuid.NewV4()
	_, err := db.Exec(`INSERT INTO address(id, street_1, street_2, city, state, zipcode) VALUES ($1, $2, $3, $4, $5, $6)`,
		addressID, address.Street1, address.Street2, address.City, address.State, address.Zipcode)
	if err != nil {
		return uuid.UUID{}, toRentError(err)
	}
	return addressID, nil
}

func addPropertyAddress(db dbConn, propertyID uuid.UUID, addressID uuid.UUID) error {
	_, err := db.Exec(`INSERT INTO property_address(property_id, address_id) VALUES ($1, $2)`, propertyID, addressID)
	if err != nil {
		return toRentError(err)
	}
	return nil
}
