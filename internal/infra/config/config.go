// Package config provides configuration management for the application.
package config

import (
	"os"
)

// Config holds all configuration for the application.
type Config struct {
	HTTPServerAddress string
	GRPCServerAddress string
}

// Load loads configuration from environment variables.
func Load() (*Config, error) {
	cfg := &Config{
		HTTPServerAddress: getEnv("EZEX_USERS_HTTP_SERVER_ADDRESS", ":8080"),
		GRPCServerAddress: getEnv("EZEX_USERS_GRPC_SERVER_ADDRESS", "0.0.0.0:50051"),
	}

	return cfg, nil
}

// BasicCheck performs basic validation of the configuration.
func (c *Config) BasicCheck() error {
	if c.HTTPServerAddress == "" {
		return nil
	}

	if c.GRPCServerAddress == "" {
		return nil
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
