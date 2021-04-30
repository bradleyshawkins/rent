package rent

import uuid "github.com/satori/go.uuid"

type Property struct {
	ID       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	Tenant   Tenant
	Landlord Landlord
	Address  Address
}

func (p *Property) ChangeTenant(t Tenant) {
	p.Tenant = t
}

func (p *Property) UpdateName(name string) {
	p.Name = name
}
