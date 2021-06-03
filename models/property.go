package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Properties map[PropertyID]*Property

type PropertyID uuid.UUID

func NewPropertyID() PropertyID {
	return PropertyID(uuid.NewV4())
}

type Property struct {
	id           PropertyID
	propertyInfo PropertyInfo
	renters      []uuid.UUID
}

type PropertyInfo struct {
	YearBuilt    time.Time
	HouseType    string
	RentPerMonth string
	PropertySize string
	LandSize     string
	Restrictions string // No Pets, No Smoking, etc
}
