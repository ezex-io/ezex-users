package port

import "github.com/ezex-io/ezex-users/internal/adapter/db/postgres/gen"

type PostgresPort interface {
	Close()
	Query() *gen.Queries
}
