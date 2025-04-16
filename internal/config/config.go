// Package config provides configuration management for the application.
package config

import (
	"fmt"

	"github.com/ezex-io/gopkg/env"
)

type Config struct {
	GRPCServerAddress string
}

func Load() (*Config, error) {
	cfg := &Config{
		GRPCServerAddress: env.GetEnv[string]("EZEX_USERS_GRPC_SERVER_ADDRESS", env.WithDefault("0.0.0.0:50051")),
	}

	return cfg, nil
}

func (c *Config) BasicCheck() error {
	if c.GRPCServerAddress == "" {
		return fmt.Errorf("GRPCServerAddress is not set")
	}

	return nil
}
