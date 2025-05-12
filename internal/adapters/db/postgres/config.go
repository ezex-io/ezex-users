package postgres

import (
	"errors"
	"fmt"
	"net"
	"net/url"
)

type Config struct {
	Host            string `koanf:"host"`
	Port            string `koanf:"ports"`
	Username        string `koanf:"username"`
	Password        string `koanf:"password"`
	Database        string `koanf:"database"`
	SSLMode         string `koanf:"sslmode"`
	TimeZone        string `koanf:"timezone"`
	MaxOpenConns    int32  `koanf:"max_open_conns"`
	MaxIdleConns    int32  `koanf:"max_idle_conns"`
	ConnMaxLifetime string `koanf:"conn_max_lifetime"`
}

func DefaultConfig() *Config {
	return &Config{
		Host:            "localhost",
		Port:            "5432",
		Username:        "postgres",
		Database:        "",
		SSLMode:         "disable",
		TimeZone:        "UTC",
		MaxOpenConns:    25,
		MaxIdleConns:    5,
		ConnMaxLifetime: "30m",
	}
}

func (c *Config) BasicCheck() error {
	if c.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (c *Config) URI() string {
	params := url.Values{}
	if c.SSLMode != "" {
		params.Add("sslmode", c.SSLMode)
	}
	if c.TimeZone != "" {
		params.Add("timezone", c.TimeZone)
	}

	hostPort := net.JoinHostPort(c.Host, c.Port)
	uri := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		url.QueryEscape(c.Username),
		url.QueryEscape(c.Password),
		hostPort,
		c.Database,
	)

	if q := params.Encode(); q != "" {
		uri += "?" + q
	}

	return uri
}
