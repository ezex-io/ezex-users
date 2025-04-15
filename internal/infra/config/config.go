package config

import (
	"fmt"
	"os"
)

type Config struct {
	HTTPServer struct {
		Address string
	}
	GRPCServer struct {
		Address string
	}
}

func Load() (*Config, error) {
	cfg := &Config{}

	cfg.HTTPServer.Address = getEnvOrDefault("EZEX_USERS_HTTP_SERVER_ADDRESS", ":8888")
	cfg.GRPCServer.Address = getEnvOrDefault("EZEX_USERS_GRPC_SERVER_ADDRESS", "0.0.0.0:50051")

	return cfg, nil
}

func (c *Config) BasicCheck() error {
	if c.HTTPServer.Address == "" {
		return fmt.Errorf("HTTP server address is required")
	}
	if c.GRPCServer.Address == "" {
		return fmt.Errorf("gRPC server address is required")
	}
	return nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
