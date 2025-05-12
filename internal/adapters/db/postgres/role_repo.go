package postgres

import (
	"context"
	"fmt"

	"github.com/ezex-io/ezex-users/internal/adapters/db/postgres/gen"
	"github.com/ezex-io/ezex-users/internal/entity/role"
	"github.com/ezex-io/ezex-users/internal/ports/repository"
)

var _ repository.RolePort = (*roleRepo)(nil)

type roleRepo struct {
	q gen.Querier
}

func newRoleRepo(q gen.Querier) repository.RolePort {
	return &roleRepo{
		q: q,
	}
}

func (r *roleRepo) GetRoleByName(ctx context.Context, name string) (*role.Role, error) {
	rol, err := r.q.GetRoleByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("GetRoleByName failed: %w", err)
	}

	return roleToDomain(&rol), nil
}
