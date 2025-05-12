package migrations

import (
	"embed"
)

//go:embed psql/*.sql
var PSQLEmbed embed.FS
