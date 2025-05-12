package postgres

import (
	"context"
	"fmt"

	"github.com/ezex-io/ezex-users/internal/adapters/db/postgres/gen"
	"github.com/ezex-io/ezex-users/internal/ports/repository"
)

var _ repository.RolePermissionsPort = (*rolePermissionsRepo)(nil)

type rolePermissionsRepo struct {
	q gen.Querier
}

func newRolePermissionsRepo(q gen.Querier) repository.RolePermissionsPort {
	return &rolePermissionsRepo{q: q}
}

func (r *rolePermissionsRepo) AddRolePermissions(ctx context.Context, roleID, permissionID string) error {
	if err := r.q.InsertRolePermission(ctx, gen.InsertRolePermissionParams{
		RoleID:       roleID,
		PermissionID: permissionID,
	}); err != nil {
		return fmt.Errorf("failed to add role permission: %w", err)
	}

	return nil
}
