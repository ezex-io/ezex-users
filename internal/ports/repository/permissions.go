package repository

import (
	"context"

	"github.com/ezex-io/ezex-users/internal/entity/permissions"
)

type PermissionsPort interface {
	GetPermission(ctx context.Context, scope, action string) (*permissions.Permission, error)
	AddPermission(ctx context.Context, name, description, scope, action string) (*permissions.Permission, error)
}
