// Package config provides configuration management for the application.
package config

import (
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
