package repository

import (
	"context"

	"github.com/ezex-io/ezex-users/internal/entity/role"
)

type RolePort interface {
	GetRoleByName(ctx context.Context, name string) (*role.Role, error)
}
