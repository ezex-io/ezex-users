// Package config provides configuration management for the application.
package config

import (
	"github.com/ezex-io/gopkg/env"
)

// Config holds all configuration for the application.
type Config struct {
	GRPCServerAddress string
}

// Load loads configuration from environment variables.
func Load() (*Config, error) {
	cfg := &Config{
		GRPCServerAddress: env.GetEnv[string]("EZEX_USERS_GRPC_SERVER_ADDRESS", env.WithDefault("0.0.0.0:50051")),
	}

	return cfg, nil
}

// BasicCheck performs basic validation of the configuration.
func (c *Config) BasicCheck() error {
	if c.GRPCServerAddress == "" {
		return nil
	}

	return nil
}
