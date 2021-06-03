package postgres

import (
	"github.com/bradleyshawkins/rent"
	uuid "github.com/satori/go.uuid"
)

func (p *Postgres) CreateLandlord() (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

func (p *Postgres) AddProperty(property rent.Property) error {
	return nil
}
