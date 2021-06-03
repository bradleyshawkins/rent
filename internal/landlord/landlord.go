package landlord

import uuid "github.com/satori/go.uuid"

type Landlord struct {
	id         uuid.UUID
	properties Properties
}

type Properties []Property

type Property struct {
	renterID uuid.UUID
}

func (l Landlord) name() {

}
