package city

import (
	"bytes"
	"encoding/csv"
	"errors"
	"github.com/google/uuid"
	cityModel "github.com/mhthrh/common-lib/model/city"
	customeError "github.com/mhthrh/common-lib/model/error"
	csvFile "github.com/mhthrh/common-lib/pkg/util/file/csv"
)

const (
	path = "customer-service/file/cities/"
	name = "cities.csv"
)

var (
	cities []cityModel.City
)

type City struct {
	path string
	name string
}

func New() cityModel.ICity {
	return City{
		path: path,
		name: name,
	}
}
func (c City) Load() *customeError.XError {
	f := csvFile.New(c.path, c.name)
	bts, e := f.Read()
	if e != nil {
		return cityModel.FileUnreachable(customeError.RunTimeError(e))
	}

	reader := csv.NewReader(bytes.NewReader(bts))

	rows, err := reader.ReadAll()
	if err != nil {
		return cityModel.FileUnreachable(customeError.RunTimeError(err))
	}
	cities = make([]cityModel.City, len(rows))
	for i, row := range rows {
		cities[i] = cityModel.City{
			ID:          uuid.New(),
			Name:        row[1],
			CountryCode: row[0],
		}
	}
	if len(cities) == 0 {
		return cityModel.FileEmpty(customeError.RunTimeError(errors.New("no city found")))
	}

	return nil
}
func (c City) Cities() ([]cityModel.City, *customeError.XError) {
	if len(cities) == 0 {
		return nil, cityModel.NotLoaded(nil)
	}
	return cities, nil
}
func (c City) GetByCountry(country string) ([]cityModel.City, *customeError.XError) {
	entry := make([]cityModel.City, 0)
	if len(cities) == 0 {
		return nil, cityModel.NotLoaded(nil)
	}
	for _, cnty := range cities {
		if cnty.CountryCode == country {
			entry = append(entry, cnty)
		}
	}
	return entry, nil
}
func (c City) GetByCity(city string) ([]cityModel.City, *customeError.XError) {
	entry := make([]cityModel.City, 0)
	if len(cities) == 0 {
		return nil, cityModel.NotLoaded(nil)
	}
	for _, cnty := range cities {
		if cnty.Name == city {
			entry = append(entry, cnty)
		}
	}
	return entry, nil
}

func (c City) GetByCityAndCountry(city, country string) (*cityModel.City, *customeError.XError) {
	if len(cities) == 0 {
		return nil, cityModel.NotLoaded(nil)
	}
	for _, cnty := range cities {
		if cnty.Name == city && cnty.CountryCode == country {
			return &cnty, nil
		}
	}
	return nil, nil
}
