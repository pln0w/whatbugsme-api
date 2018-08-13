package env

import (
	"os"
)

type Config struct {
	Environment string
	APITitle    string
	APIVersion  string
	APIPort     string
	Server      string
	Database    string
}

// Get returns Config struct pointer
func Get() *Config {
	return &Config{
		Environment: os.Getenv("ENV"),
		APITitle:    os.Getenv("API_TITLE"),
		APIVersion:  os.Getenv("API_VERSION"),
		APIPort:     os.Getenv("API_PORT"),
		Server:      os.Getenv("SERVER"),
		Database:    os.Getenv("DATABASE"),
	}
}
