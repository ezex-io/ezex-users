package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	dbPass := "secret"
	require.NoError(t, os.Setenv("EZEX_USERS_DATABASE_PASSWORD", dbPass))

	cfg, err := Load("./config.example.toml")
	require.NoError(t, err)
	require.NoError(t, cfg.BasicCheck())

	require.NotNil(t, cfg)
	require.NotNil(t, cfg.Database)
	assert.Equal(t, dbPass, cfg.Database.Password)
}
