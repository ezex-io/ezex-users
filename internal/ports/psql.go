package ports

import (
	"context"

	"github.com/ezex-io/ezex-users/internal/ports/repository"
)

type PostgresPort interface {
	Close()
	MigrateUp(ctx context.Context, email, username, password string) error
	MigrateDown() error

	PostgresRepositories
}

type PostgresRepositories interface {
	User() repository.UserPort
	Role() repository.RolePort
	Permissions() repository.PermissionsPort
	RolePermissions() repository.RolePermissionsPort
	SecurityImage() repository.SecurityImagePort
}
