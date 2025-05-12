package postgres

import (
	"context"
	"fmt"

	"github.com/ezex-io/ezex-users/internal/adapters/db/postgres/gen"
	"github.com/ezex-io/ezex-users/internal/entity/permissions"
	"github.com/ezex-io/ezex-users/internal/ports/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

var _ repository.PermissionsPort = (*permissionsRepo)(nil)

type permissionsRepo struct {
	q gen.Querier
}

func newPermissionsRepo(q gen.Querier) repository.PermissionsPort {
	return &permissionsRepo{
		q: q,
	}
}

func (p *permissionsRepo) GetPermission(ctx context.Context, scope, action string) (*permissions.Permission, error) {
	perm, err := p.q.GetPermissionByScopeAction(ctx, gen.GetPermissionByScopeActionParams{
		Scope:  scope,
		Action: action,
	})
	if err != nil {
		return nil, fmt.Errorf("GetPermission failed for %s:%s: %w", scope, action, err)
	}

	return permissionToDomain(&perm), nil
}

func (p *permissionsRepo) AddPermission(ctx context.Context, name, description,
	scope, action string,
) (*permissions.Permission, error) {
	perm, err := p.q.UpsertPermission(ctx, gen.UpsertPermissionParams{
		Name:        pgtype.Text{String: name, Valid: true},
		Description: pgtype.Text{String: description, Valid: true},
		Scope:       scope,
		Action:      action,
	})
	if err != nil {
		return nil, fmt.Errorf("AddPermission failed for %s:%s: %w", name, description, err)
	}

	return permissionToDomain(&perm), nil
}
