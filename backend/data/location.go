package data

type Address struct {
	Street  string `json:"street"`
	Street2 string `json:"street2"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

type GeoLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type DistanceUnit string

const (
	Miles      DistanceUnit = "mi"
	Kilometers DistanceUnit = "km"
)
