package config

import (
	"fmt"
	"github.com/alhaos/quick-menu-api/internal/database"
	"github.com/ilyakaznacheev/cleanenv"
)

type Configuration struct {
	Database database.Config `yaml:"database"`
}

// New pointer to Configuration
func New(filename string) (*Configuration, error) {

	c := &Configuration{}

	err := cleanenv.ReadConfig(filename, c)
	if err != nil {
		return nil, fmt.Errorf("error reading configuration: %w", err)
	}

	return c, nil
}
