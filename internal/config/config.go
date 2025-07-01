package config

import (
	"fmt"
	"github.com/alhaos/quick-menu-api/internal/authService"
	"github.com/alhaos/quick-menu-api/internal/database"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Configuration struct {
	Database    database.Config `yaml:"database"`
	Address     Address         `yaml:"address"`
	AuthService authService.Config
}

type Address struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
}

// New creates and returns a new Configuration instance by loading settings from
// both a configuration file and environment variables.
//
// Parameters:
//   - filename: path to the YAML configuration file
//
// Returns:
//   - *Configuration: pointer to the loaded configuration
//   - error: error if configuration fails to load or required environment variables are missing
//
// Configuration is loaded in the following order:
//  1. Base parameters from the YAML file
//  2. Environment variables (for fields with env tag)
//
// Required environment variables:
//   - QUICK_MENU_SECRET: application secret key (must be set)
//
// Example usage:
//
//	cfg, err := config.New("config.yml")
//	if err != nil {
//	    log.Fatalf("Failed to load configuration: %v", err)
//	}
//
// Note:
//   - The Secret field is intentionally excluded from YAML (yaml:"-")
//     and must be provided via environment variable
func New(filename string) (*Configuration, error) {

	c := &Configuration{}

	err := cleanenv.ReadConfig(filename, c)
	if err != nil {
		return nil, fmt.Errorf("error reading configuration: %w", err)
	}

	secret := os.Getenv("QUICK_MENU_SECRET")
	if secret == "" {
		return nil, fmt.Errorf("environment variable QUICK_MENU_SECRET is not set")
	}
	c.AuthService.Secret = []byte(secret)
	return c, nil
}
