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

func (p *Postgres) RemoveProperty(accountID, propertyID uuid.UUID) error {
	tx, err := p.db.Begin()
	if err != nil {
		return toRentError(err)
	}

	prop, err := loadProperty(tx, accountID, propertyID)
	if err != nil {
		return err
	}

	err = prop.Disable()
	if err != nil {
		return err
	}

	err = updatePropertyStatus(tx, prop)
	if err != nil {
		return err
	}

	err = removePropertyAddress(tx, prop.ID, prop.Address.ID)
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

func loadProperty(db dbConn, accountID, propertyID uuid.UUID) (*rent.Property, error) {
	var propID, addrID uuid.UUID
	var name, street1, street2, city, state, zipcode string
	var statusID rent.PropertyStatus
	err := db.QueryRow(`SELECT p.id, p.name, p.property_status_id, a.id, a.street_1, a.street_2, a.city, a.state, a.zipcode
								FROM property p
								INNER JOIN property_address pa ON p.id = pa.property_id
								INNER JOIN address a ON pa.address_id = a.id
								WHERE p.id = $1 AND p.account_id = $2`, propertyID, accountID).
		Scan(&propID, &name, &statusID, &addrID, &street1, &street2, &city, &state, &zipcode)
	if err != nil {
		return nil, toRentError(err)
	}
	addr, err := rent.NewExistingAddress(addrID, street1, street2, city, state, zipcode)
	if err != nil {
		return nil, err
	}
	return rent.NewExistingProperty(propID, name, statusID, addr)
}

func updatePropertyStatus(db dbConn, property *rent.Property) error {
	_, err := db.Exec(`UPDATE property SET property_status_id = $1 WHERE id = $2`, property.Status, property.ID)
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

func removePropertyAddress(db dbConn, propertyID, addressID uuid.UUID) error {
	_, err := db.Exec(`DELETE FROM property_address WHERE property_id = $1 and address_id = $2`, propertyID, addressID)
	if err != nil {
		return toRentError(err)
	}
	return nil
}
