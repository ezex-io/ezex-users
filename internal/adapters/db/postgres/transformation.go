package postgres

import (
	"github.com/ezex-io/ezex-users/internal/adapters/db/postgres/gen"
	"github.com/ezex-io/ezex-users/internal/entity/permissions"
	"github.com/ezex-io/ezex-users/internal/entity/role"
	"github.com/ezex-io/ezex-users/internal/entity/user"
)

func userToDomain(_ *gen.User) *user.User {
	return &user.User{}
}

func permissionToDomain(perm *gen.Permission) *permissions.Permission {
	return &permissions.Permission{
		ID:          perm.ID,
		Name:        perm.Name.String,
		Description: perm.Description.String,
		Scope:       perm.Scope,
		Action:      perm.Action,
	}
}

func roleToDomain(rol *gen.Role) *role.Role {
	return &role.Role{
		ID:          rol.ID,
		Name:        rol.Name,
		CreatedAt:   rol.CreatedAt,
		UpdatedAt:   rol.UpdatedAt,
		DeletedAt:   rol.DeletedAt.Time,
		CreatedByID: rol.CreatedByID.String,
		UpdatedByID: rol.UpdatedByID.String,
		DeletedByID: rol.DeletedByID.String,
		IsSystem:    rol.IsSystem,
		IsDefault:   rol.IsDefault,
	}
}
