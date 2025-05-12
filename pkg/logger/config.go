package logger

type Config struct {
	Colorful           bool              `koanf:"colorful"`
	MaxBackups         int               `koanf:"max_backups"`
	RotateLogAfterDays int               `koanf:"rotate_log_after_days"`
	Compress           bool              `koanf:"compress"`
	Targets            []string          `koanf:"targets"`
	Levels             map[string]string `koanf:"levels"`
}

func DefaultConfig() *Config {
	conf := &Config{
		Levels:             make(map[string]string),
		Colorful:           true,
		MaxBackups:         0,
		RotateLogAfterDays: 1,
		Compress:           true,
		Targets:            []string{"console"},
	}

	conf.Levels["default"] = "info"
	conf.Levels["_grpc"] = "info"
	conf.Levels["_interactor"] = "warn"
	conf.Levels["_database"] = "warn"
	conf.Levels["_cache"] = "warn"
	conf.Levels["_migration"] = "info"

	return conf
}

// BasicCheck performs basic checks on the configuration.
func (*Config) BasicCheck() error {
	return nil
}
