package rent_test

import (
	"testing"

	"github.com/bradleyshawkins/rent"
	"github.com/matryer/is"
	uuid "github.com/satori/go.uuid"
)

func TestNewProperty(t *testing.T) {
	i := is.New(t)
	addr, err := rent.NewAddress("street1", "street2", "city", "state", "zipcode")
	i.NoErr(err)

	prop, err := rent.NewProperty("test property", addr)
	i.NoErr(err)

	i.True(prop.ID != (uuid.UUID{}))
	i.Equal(prop.Name, "test property")
	i.Equal(prop.Status, rent.PropertyVacant)
	i.Equal(prop.Address.Street1, "street1")
	i.Equal(prop.Address.Street2, "street2")
	i.Equal(prop.Address.City, "city")
	i.Equal(prop.Address.State, "state")
	i.Equal(prop.Address.Zipcode, "zipcode")
}

func TestNewProperty_MissingName(t *testing.T) {
	i := is.New(t)
	addr, err := rent.NewAddress("street1", "street2", "city", "state", "zipcode")
	i.NoErr(err)

	_, err = rent.NewProperty("", addr)
	i.True(err != nil)

	e, ok := err.(*rent.Error)
	i.True(ok)

	i.True(len(e.InvalidFields()) == 1)

	i.True(e.InvalidFields()[0].Reason == rent.ReasonMissing)
}

func TestNewAddress_MissingField(t *testing.T) {
	street1 := "street1"
	street2 := "street2"
	city := "city"
	state := "state"
	zipcode := "zipcode"

	tests := []struct {
		name    string
		street1 string
		street2 string
		city    string
		state   string
		zipcode string
	}{
		{
			name:    "Missing Street1",
			street1: "", street2: street2, city: city, state: state, zipcode: zipcode,
		},
		{
			name:    "Missing city",
			street1: street1, street2: street2, city: "", state: state, zipcode: zipcode,
		},
		{
			name:    "Missing Street1",
			street1: street1, street2: street2, city: city, state: "", zipcode: zipcode,
		},
		{
			name:    "Missing Zipcode",
			street1: street1, street2: street2, city: city, state: state, zipcode: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := is.New(t)
			_, err := rent.NewAddress(tt.street1, tt.street2, tt.city, tt.state, tt.zipcode)
			i.True(err != nil)

			e, ok := err.(*rent.Error)
			i.True(ok)

			i.True(len(e.InvalidFields()) == 1)

			i.True(e.InvalidFields()[0].Reason == rent.ReasonMissing)
		})
	}
}
