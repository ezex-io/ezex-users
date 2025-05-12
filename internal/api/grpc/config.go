package grpc

type Config struct {
	Address          string `koanf:"address"`
	Port             int    `koanf:"ports"`
	EnableReflection bool   `koanf:"enable_reflection"`
	EnableHealth     bool   `koanf:"enable_health"`
}

func DefaultConfig() *Config {
	return &Config{
		Address: "0.0.0.0",
		Port:    50051,
	}
}

func (*Config) BasicCheck() error {
	return nil
}
