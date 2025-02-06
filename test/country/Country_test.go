package country_test

import (
	"customer-service/pkg/address/country"
	customeError "github.com/mhthrh/common-lib/errors"
	"github.com/mhthrh/common-lib/model/test"
	"testing"
)

var (
	c *country.Countries1
	e *customeError.XError
)

func init() {
	c, e = country.LoadCountries()
}

func TestLoadCountry(t *testing.T) {

	if e != nil {
		t.Error(e)
	}
	if len(c.Countries) < 10 {
		t.Error("countries length should be greater than 10")
	}
	if c.Countries[0].Code != "AF" {
		t.Error("code should be AF")
	}
}

func TestFilterByCode(t *testing.T) {

	tests := []test.Test{
		{
			Name:     "with valid code",
			Input:    []string{"IR"},
			OutPut:   1,
			HasError: false,
			Err:      nil,
		},
		{
			Name:     "with invalid code",
			Input:    []string{"XNXX", "uweyuf"},
			OutPut:   0,
			HasError: false,
			Err:      nil,
		},
		{
			Name:     "without any codes",
			Input:    []string{},
			OutPut:   0,
			HasError: false,
			Err:      nil,
		},
	}
	for _, tst := range tests {

		cnt, e := c.FilterByCode(tst.Input.([]string)...)
		if e != nil {
			t.Error(tst.Err)
		}
		if tst.HasError {
			if e.Code != tst.Err.Code {
				t.Error("expected error is not equal with this error.")
			}
			break
		}
		if len(cnt.Countries) != tst.OutPut.(int) {
			t.Errorf("country count length should be %d", tst.OutPut.(int))
		}
	}
}
