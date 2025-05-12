// Package config provides configuration management for the application.
package config

import (
	"strings"

	"github.com/ezex-io/ezex-users/internal/adapters/db/postgres"
	"github.com/ezex-io/ezex-users/internal/api/grpc"
	"github.com/ezex-io/ezex-users/pkg/logger"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Development bool             `koanf:"development"`
	GRPC        *grpc.Config     `koanf:"grpc"`
	Database    *postgres.Config `koanf:"database"`
	Logger      *logger.Config   `koanf:"logger"`
}

func Load(path string) (*Config, error) {
	cfg := _default()

	koan := koanf.New(".")

	if err := koan.Load(file.Provider(path), toml.Parser()); err != nil {
		return nil, err
	}

	// Load from env (after file)
	if err := koan.Load(env.Provider("EZEX_USERS_", ".", func(key string) string {
		key = strings.TrimPrefix(key, "EZEX_USERS_")
		key = strings.ToLower(key)
		parts := strings.Split(key, "_")
		if len(parts) > 1 {
			return parts[0] + "." + strings.Join(parts[1:], "_")
		}

		return key
	}), nil); err != nil {
		return nil, err
	}

	if err := koan.Unmarshal("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) BasicCheck() error {
	if err := c.GRPC.BasicCheck(); err != nil {
		return err
	}

	if err := c.Database.BasicCheck(); err != nil {
		return err
	}

	return c.Logger.BasicCheck()
}

func _default() *Config {
	return &Config{
		GRPC:     grpc.DefaultConfig(),
		Database: postgres.DefaultConfig(),
		Logger:   logger.DefaultConfig(),
	}
}
