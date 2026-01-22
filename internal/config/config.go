// Package config contains configuration structures and loaders for the auth service.
package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config describes the root configuration of the auth service.
//
// It aggregates configuration sections for all major subsystems,
// including HTTP server, database connections, cache, and
// authentication settings.
type Config struct {
	// Env defines the current application environment
	// (for example: dev, prod).
	Env string `yaml:"env" env:"ENV" env-default:"dev"`

	// HTTP contains configuration for the HTTP server.
	HTTP HTTPConfig `yaml:"http"`

	// Postgres contains PostgreSQL connection configuration.
	Postgres PostgresConfig `yaml:"postgres"`

	// Redis contains Redis connection configuration.
	Redis RedisConfig `yaml:"redis"`

	// Auth contains authentication-related configuration.
	Auth AuthConfig `yaml:"auth"`
}

// Load reads configuration from the provided file path
// and overrides values using environment variables.
func Load(path string) (*Config, error) {
	if path == "" {
		return nil, fmt.Errorf("config path is empty")
	}

	if _, err := os.Stat(path); err != nil {
		return nil, fmt.Errorf("config file not found: %w", err)
	}

	var cfg Config

	// Read YAML configuration file.
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	// Override configuration values from environment variables.
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("read env: %w", err)
	}

	return &cfg, nil
}

// MustLoad is a helper that wraps Load and panics on error.
func MustLoad(path string) *Config {
	cfg, err := Load(path)
	if err != nil {
		panic(err)
	}

	return cfg
}
