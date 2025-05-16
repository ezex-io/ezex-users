package grpc

import "github.com/ezex-io/gopkg/env"

type Config struct {
	Address           string
	EnableHealthCheck bool
	// EnableReflection enable reflection option in grpc to load rpc method on client (postman) without proto
	EnableReflection bool
}

func LoadFromEnv() *Config {
	cfg := &Config{
		Address:           env.GetEnv[string]("EZEX_USERS_GRPC_ADDRESS", env.WithDefault("0.0.0.0:50051")),
		EnableHealthCheck: env.GetEnv[bool]("EZEX_USERS_GRPC_ENABLE_HEALTH_CHECK", env.WithDefault("false")),
		EnableReflection:  env.GetEnv[bool]("EZEX_USERS_GRPC_ENABLE_REFLECTION", env.WithDefault("false")),
	}

	return cfg
}

func (*Config) BasicCheck() error {
	return nil
}
