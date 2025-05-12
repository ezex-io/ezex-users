package apps

import (
	_ "embed"

	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"
)

//go:embed service.toml
var svcBuf []byte

type Service struct {
	Name        string       `koanf:"name"`
	Description string       `koanf:"description"`
	Permissions []Permission `koanf:"permissions"`
}

type Permission struct {
	Name        string `koanf:"name"`
	Description string `koanf:"description"`
	Scope       string `koanf:"scope"`
	Action      string `koanf:"action"`
}

func load() (*Service, error) {
	koan := koanf.New(".")
	if err := koan.Load(rawbytes.Provider(svcBuf), toml.Parser()); err != nil {
		return nil, err
	}

	var svc Service
	if err := koan.Unmarshal("", &svc); err != nil {
		return nil, err
	}

	return &svc, nil
}
