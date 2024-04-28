package models

import (
	"fmt"
)

type Address struct {
	Street  string `json:"street,omitempty"`
	Street2 string `json:"street2,omitempty"`
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	ZipCode string `json:"zipCode,omitempty"`
}

// FullAddress returns the full address as a single string
func (a Address) FullAddress() string {
	return fmt.Sprintf("%s %s, %s, %s, %s", a.Street, a.Street2, a.City, a.State, a.ZipCode)
}

// CityStateZip returns the city, state, and zip code as a single string
func (a Address) CityStateZip() string {
	return fmt.Sprintf("%s, %s, %s", a.City, a.State, a.ZipCode)
}
