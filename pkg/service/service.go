package service

import (
	"errors"
	modelAddress "github.com/mhthrh/common-lib/model/address"
	modelError "github.com/mhthrh/common-lib/model/error"
	"strings"
)

type Address struct {
	Adrs modelAddress.Address
}

func NewAddress(street, postalCode, state, cntry, cty string) (*Address, *modelError.XError) {
	if strings.Trim(street, " ") == "" {
		return nil, modelAddress.StreetNotFound(nil)
	}
	if strings.Trim(postalCode, " ") == "" {
		return nil, modelAddress.PostalCodeNotFound(nil)
	}
	if strings.Trim(state, " ") == "" {
		return nil, modelAddress.StateNotFound(nil)
	}
	if strings.Trim(cntry, " ") == "" {
		return nil, modelAddress.CountryNotFound(nil)
	}
	if strings.Trim(cty, " ") == "" {
		return nil, modelAddress.CityNotFound(nil)
	}
	c := country.New()
	err := c.Load()
	if err != nil {
		return nil, modelAddress.CountryNotFound(err)
	}
	cResult, err := c.GetByCode(cntry)
	if err != nil {
		return nil, modelAddress.CityNotFound(err)
	}
	if cResult == nil {
		return nil, modelAddress.CountryNotFound(modelError.RunTimeError(errors.New("invalid countries length")))
	}
	ci := city.New()
	err = ci.Load()
	if err != nil {
		return nil, modelAddress.CityNotFound(err)
	}
	cyResult, err := ci.GetByCityAndCountry(cty, cntry)

	if err != nil {
		return nil, modelAddress.CityNotFound(modelError.RunTimeError(errors.New("count is more than 1")))
	}

	return &Address{Adrs: Address{
		Street:     strings.Trim(street, " "),
		City:       cyResult.Cities[0],
		State:      state,
		PostalCode: postalCode,
		Country:    cResult.Countries[0],
	}}, nil
}
