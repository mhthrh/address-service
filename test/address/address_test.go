package address

import (
	"customer-service/pkg/address"
	customeError "github.com/mhthrh/common-lib/errors/address"
	"github.com/mhthrh/common-lib/model/test"
	"testing"
)

func TestNewAddress(t *testing.T) {
	type adTest struct {
		street     string
		postalCode string
		state      string
		country    string
		city       string
	}
	tests := []test.Test{
		{
			Name: "address test 1",
			Input: adTest{
				street:     "",
				postalCode: "123",
				state:      "state",
				country:    "country",
				city:       "city",
			},
			OutPut:   nil,
			HasError: true,
			Err:      customeError.StreetNotFound(nil),
		}, {
			Name: "address test 2",
			Input: adTest{
				street:     "whiting way",
				postalCode: "",
				state:      "state",
				country:    "country",
				city:       "city",
			},
			OutPut:   nil,
			HasError: true,
			Err:      customeError.PostalCodeNotFound(nil),
		}, {
			Name: "address test 3",
			Input: adTest{
				street:     "whiting way",
				postalCode: "123",
				state:      "",
				country:    "country",
				city:       "city",
			},
			OutPut:   nil,
			HasError: true,
			Err:      customeError.StateNotFound(nil),
		}, {
			Name: "address test 4",
			Input: adTest{
				street:     "whiting way",
				postalCode: "123",
				state:      "state",
				country:    "",
				city:       "city",
			},
			OutPut:   nil,
			HasError: true,
			Err:      customeError.CountryNotFound(nil),
		}, {
			Name: "address test 5",
			Input: adTest{
				street:     "whiting way",
				postalCode: "123",
				state:      "state",
				country:    "country",
				city:       "",
			},
			OutPut:   nil,
			HasError: true,
			Err:      customeError.CityNotFound(nil),
		}, {
			Name: "address test 6",
			Input: adTest{
				street:     "whiting way",
				postalCode: "123",
				state:      "state",
				country:    "LLREE",
				city:       "city",
			},
			OutPut:   nil,
			HasError: true,
			Err:      customeError.CountryNotFound(nil),
		}, {
			Name: "address test 7",
			Input: adTest{
				street:     "whiting way",
				postalCode: "123",
				state:      "state",
				country:    "GB",
				city:       "London",
			},
			OutPut:   nil,
			HasError: false,
			Err:      nil,
		},
	}
	for _, tst := range tests {
		_, e := address.NewAddress(tst.Input.(adTest).street, tst.Input.(adTest).postalCode, tst.Input.(adTest).state, tst.Input.(adTest).country, tst.Input.(adTest).city)
		if tst.HasError {
			if e.Code != tst.Err.Code {
				t.Errorf(`teast name %s expected error "%v" but got "%v"`, tst.Name, tst.Err, e)
				break
			}
		}
		if e != nil && !tst.HasError {
			t.Errorf(` test name %s expected "%v" but got "%v"`, tst.Name, tst.Err, e)
		}
	}

}
