package postgres

import "github.com/bradleyshawkins/rent"

func (p *Postgres) RegisterRenter(r rent.Renter) error {
	return nil
}

func (p *Postgres) ApplyForProperty(r *rent.Renter, pr *rent.Property) error {
	return nil
}
