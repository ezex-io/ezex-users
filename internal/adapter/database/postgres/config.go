package postgres

import (
	"fmt"
	"net/url"

	"github.com/ezex-io/gopkg/env"
)

type Config struct {
	Address  string
	Database string
	Username string
	Password string

	MaxOpenConns    int32
	MaxIdleConns    int32
	ConnMaxLifetime string
}

func LoadFromEnv() *Config {
	cfg := &Config{
		Address:         env.GetEnv[string]("EZEX_USERS_DB_ADDRESS", env.WithDefault("0.0.0.0:5432")),
		Database:        env.GetEnv[string]("EZEX_USERS_DB_DATABASE", env.WithDefault("ezex_users")),
		Username:        env.GetEnv[string]("EZEX_USERS_DB_USERNAME"),
		Password:        env.GetEnv[string]("EZEX_USERS_DB_PASSWORD"),
		MaxOpenConns:    25,
		MaxIdleConns:    5,
		ConnMaxLifetime: "30m",
	}

	return cfg
}

func (*Config) BasicCheck() error {
	return nil
}

func (c *Config) uri() string {
	params := url.Values{}
	uri := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		url.QueryEscape(c.Username),
		url.QueryEscape(c.Password),
		c.Address,
		c.Database,
	)

	if q := params.Encode(); q != "" {
		uri += "?" + q
	}

	return uri
}
