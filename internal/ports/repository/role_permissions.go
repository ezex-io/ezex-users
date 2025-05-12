package repository

import (
	"context"
)

type RolePermissionsPort interface {
	AddRolePermissions(ctx context.Context, roleID, permissionID string) error
}
