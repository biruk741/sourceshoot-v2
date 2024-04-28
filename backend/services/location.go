package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"

	config2 "backend/config"
	serviceTypes "backend/services/types"
)

//go:generate mockgen -destination=mock_locationService.go -package=locationService . LocationService

type LocationService interface {
	GetCoordinatesFromAddress(location string) (serviceTypes.GeoLocation, error)
	GetLocationFromCoordinates(coords serviceTypes.GeoLocation) (string, error)
}

type LocationInstance struct {
	GeocodingAPIURL string
	GeocodingAPIKey string
}

func NewLocationService() (LocationService, error) {
	config, err := config2.LoadConfig()
	if err != nil {
		return LocationInstance{}, err
	}
	return LocationInstance{
		GeocodingAPIURL: config.GeocodingAPIURL,
		GeocodingAPIKey: config.GeocodingAPIKey,
	}, nil
}

// GoogleGeocodeResponse represents the response structure from the Google Geocoding API.
type GoogleGeocodeResponse struct {
	Results []struct {
		Geometry struct {
			Location serviceTypes.GeoLocation `json:"location"`
		} `json:"geometry"`
		FormattedAddress string `json:"formatted_address"`
	} `json:"results"`
	Status string `json:"status"`
}

// GetCoordinates fetches a GeoLocation for a given address or location name.
func (s LocationInstance) GetCoordinatesFromAddress(address string) (serviceTypes.GeoLocation, error) {
	var loc serviceTypes.GeoLocation
	resp, err := http.Get(fmt.Sprintf("%s?address=%s&key=%s", s.GeocodingAPIURL, address, s.GeocodingAPIKey))
	if err != nil {
		return loc, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return loc, errors.New("failed to fetch geocoding data")
	}

	var geocodeResponse GoogleGeocodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&geocodeResponse); err != nil {
		return loc, err
	}

	if geocodeResponse.Status != "OK" || len(geocodeResponse.Results) == 0 {
		return loc, errors.New("invalid geocoding response")
	}

	return geocodeResponse.Results[0].Geometry.Location, nil
}

// GetLocation fetches an address for given GeoLocation coordinates using the Google Reverse Geocoding API.
func (s LocationInstance) GetLocationFromCoordinates(coords serviceTypes.GeoLocation) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s?latlng=%f,%f&key=%s", s.GeocodingAPIURL, coords.Latitude, coords.Longitude, s.GeocodingAPIKey))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to fetch reverse geocoding data")
	}

	var geocodeResponse GoogleGeocodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&geocodeResponse); err != nil {
		return "", err
	}

	if geocodeResponse.Status != "OK" || len(geocodeResponse.Results) == 0 {
		return "", errors.New("invalid reverse geocoding response")
	}

	return geocodeResponse.Results[0].FormattedAddress, nil
}

// CalculateDistance computes the distance between two GeoLocation points.
// It uses the "haversine" formula to compute the great-circle distance between the points.
func CalculateDistance(loc1, loc2 serviceTypes.GeoLocation) float64 {
	const R = 6371 // Earth radius in kilometers

	dLat := (loc2.Latitude - loc1.Latitude) * (math.Pi / 180.0)
	dLon := (loc2.Longitude - loc1.Longitude) * (math.Pi / 180.0)

	a := (math.Sin(dLat/2) * math.Sin(dLat/2)) +
		(math.Cos(loc1.Latitude*(math.Pi/180.0)) * math.Cos(loc2.Latitude*(math.Pi/180.0)) * math.Sin(dLon/2) * math.Sin(dLon/2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c // Returns distance in kilometers
}
