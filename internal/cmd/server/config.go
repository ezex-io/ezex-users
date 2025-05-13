package main

import (
	"github.com/ezex-io/ezex-users/internal/adapter/db/postgres"
	"github.com/ezex-io/ezex-users/internal/adapter/grpc"
)

type Config struct {
	GRPC     *grpc.Config
	Database *postgres.Config
}

func makeConfig() *Config {
	return &Config{
		GRPC:     grpc.LoadFromEnv(),
		Database: postgres.LoadFromEnv(),
	}
}

func (c *Config) BasicCheck() error {
	if err := c.GRPC.BasicCheck(); err != nil {
		return err
	}

	return c.Database.BasicCheck()
}
