package rent

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Property struct {
	id                uuid.UUID
	propertyInfo      PropertyInfo
	currentRenter     *Renter
	approvedApplicant *Renter
	applicants        map[uuid.UUID]*Renter
}

type PropertyInfo struct {
	YearBuilt    time.Time
	HouseType    string
	RentPerMonth string
	HomeSize     string
	LandSize     string
	Restrictions string // No Pets, No Smoking, etc
}

type Properties map[uuid.UUID]*Property

func NewEmptyProperties() Properties {
	return make(map[uuid.UUID]*Property)
}
