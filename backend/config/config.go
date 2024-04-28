package config

import (
	"fmt"
	_ "log"
	"os"

	"github.com/joho/godotenv"
)

// Config represents the configuration variables.
type Config struct {
	Environment string

	ServerPort    string
	ServerAddress string

	DbName     string
	DBUsername string
	DBPassword string
	DBHostname string
	DBPort     string

	WebsiteAddress       string
	WebsiteAddressRemote string
	WebsitePort          string

	StoryBookAddress string
	StoryBookPort    string

	FirebaseConfigDir string

	GeocodingAPIURL string
	GeocodingAPIKey string

	RateLimiterRequestsPerSecond string
	DisableAuthMiddleware        string
	MockFirebaseID               string
}

// LoadConfig loads the configuration variables from the .env file.
func LoadConfig() (*Config, error) {
	// Load environment variables from file
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	// Create a new Config struct and populate it with the environment variables
	config := &Config{
		Environment:   os.Getenv("ENVIRONMENT"),
		ServerAddress: os.Getenv("SERVER"),
		ServerPort:    os.Getenv("PORT"),

		RateLimiterRequestsPerSecond: os.Getenv("RATE_LIMITER_REQUESTS_PER_SECOND"),

		DbName:     os.Getenv("DB_NAME"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHostname: os.Getenv("DB_HOSTNAME"),
		DBPort:     os.Getenv("DB_PORT"),

		WebsiteAddress:        os.Getenv("WEBSITE_ADDRESS"),
		WebsiteAddressRemote:  os.Getenv("WEBSITE_ADDRESS_REMOTE"),
		WebsitePort:           os.Getenv("WEBSITE_PORT"),
		StoryBookAddress:      os.Getenv("STORYBOOK_ADDRESS"),
		StoryBookPort:         os.Getenv("STORYBOOK_PORT"),
		FirebaseConfigDir:     os.Getenv("FIREBASE_CONFIG_DIR"),
		GeocodingAPIURL:       os.Getenv("GEOCODING_API_URL"),
		GeocodingAPIKey:       os.Getenv("GEOCODING_API_KEY"),
		DisableAuthMiddleware: os.Getenv("DISABLE_AUTH_MIDDLEWARE"),
		MockFirebaseID:        os.Getenv("MOCK_FIREBASE_ID"),
	}

	return config, nil
}
