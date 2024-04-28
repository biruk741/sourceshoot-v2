package serviceTypes

import (
	"backend/data"
)

type GeoLocation = data.GeoLocation
type Address = data.Address
type DistanceUnit = data.DistanceUnit

type Service struct {
	GeocodingAPIURL string
	GeocodingAPIKey string
}
