package city_test

import (
	"customer-service/pkg/address/city"
	cityModel "github.com/mhthrh/common-lib/model/address/city"
	"github.com/mhthrh/common-lib/model/test"
	"testing"
)

var (
	c cityModel.ICity
)

func init() {
	c = city.New()
}
func TestLoad(t *testing.T) {
	err := c.Load()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
func TestCities(t *testing.T) {
	tests := []test.Test{
		{
			Name:     "get all cities",
			Input:    nil,
			OutPut:   47868,
			HasError: false,
			Err:      nil,
		},
	}

	for _, tst := range tests {
		cities, err := c.Cities()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if len(cities) != tst.OutPut.(int) {
			t.Errorf("test name %s, expected %d, got %d", tst.Name, tst.OutPut.(int), len(cities))
		}
	}
}

func TestGetByCountry(t *testing.T) {
	tests := []test.Test{
		{
			Name:     "filter by code",
			Input:    "IR",
			OutPut:   491,
			HasError: false,
			Err:      nil,
		}, {
			Name:     "filter by wrong code",
			Input:    "IROO",
			OutPut:   0,
			HasError: false,
			Err:      nil,
		},
	}
	for _, tst := range tests {
		result, err := c.GetByCountry(tst.Input.(string))
		if err != nil && !tst.HasError {
			t.Errorf("expected no error, got %v", err)
		}
		if len(result) != tst.OutPut.(int) {
			t.Errorf("test name %s, expected %d, got %d", tst.Name, tst.OutPut.(int), len(result))
		}

	}
}

func TestGetByCity(t *testing.T) {
	tests := []test.Test{
		{
			Name:     "filter by city",
			Input:    "London",
			OutPut:   3,
			HasError: false,
			Err:      nil,
		}, {
			Name:     "filter by wrong city",
			Input:    "WXYZ",
			OutPut:   0,
			HasError: false,
			Err:      nil,
		},
	}
	for _, tst := range tests {
		result, err := c.GetByCity(tst.Input.(string))
		if err != nil && !tst.HasError {
			t.Errorf("expected no error, got %v", err)
		}
		if len(result) != tst.OutPut.(int) {
			t.Errorf("test name %s, expected %d, got %d", tst.Name, tst.OutPut.(int), len(result))
		}
	}
}

func TestGetByCityAndCountry(t *testing.T) {
	tests := []test.Test{
		{
			Name:     "filter by city and country",
			Input:    []string{"London", "GB"},
			OutPut:   1,
			HasError: false,
			Err:      nil,
		}, {
			Name:     "filter by wrong city and country",
			Input:    []string{"GBXX", "XLondonX"},
			OutPut:   0,
			HasError: false,
			Err:      nil,
		},
	}
	for i, tst := range tests {
		result, err := c.GetByCityAndCountry(tst.Input.([]string)[0], tst.Input.([]string)[1])
		if err != nil && !tst.HasError {
			t.Errorf("expected no error, got %v", err)
		}
		switch i {
		case 0:
			if result.CountryCode != tst.Input.([]string)[0] && result.CountryCode != tst.Input.([]string)[1] {
				t.Errorf("test name %s expected %s, got %s", tst.Name, tst.Input.([]string)[0], result.CountryCode)
			}
		case 1:
			if result != nil {
				t.Errorf("test name %s expected nil, got %s", tst.Name, tst.Input.([]string)[0])
			}
		}
	}
}
