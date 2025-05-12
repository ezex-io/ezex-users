package apps

import (
	"context"

	"github.com/ezex-io/ezex-users/internal/adapters/db/postgres"
	"github.com/ezex-io/ezex-users/internal/api/grpc"
	"github.com/ezex-io/ezex-users/internal/config"
	"github.com/ezex-io/ezex-users/internal/interactors/user"
	"github.com/ezex-io/ezex-users/internal/ports"
	"github.com/ezex-io/ezex-users/pkg/logger"
)

type Users struct {
	Config  *config.Config
	Service *Service

	grpcServer *grpc.Server

	psql ports.PostgresPort
}

func NewUsers(ctx context.Context, cfg *config.Config, init bool) (*Users, error) {
	users := &Users{
		Config: cfg,
	}

	logger.InitGlobalLogger(cfg.Logger)

	svc, err := load()
	if err != nil {
		return nil, err
	}

	users.Service = svc

	psql, err := postgres.New(ctx, cfg.Database)
	if err != nil {
		return nil, err
	}

	logger.Debug("initializing database",
		"host", cfg.Database.Host,
		"ports", cfg.Database.Port,
		"database", cfg.Database.Database,
	)

	users.psql = psql

	if !init {
		userInteractor := user.New(psql.SecurityImage())

		grpcSrv := grpc.New(cfg.GRPC, userInteractor)
		users.grpcServer = grpcSrv
	}

	return users, nil
}

func (a *Users) Init(ctx context.Context) error {
	return a.syncPermissions(ctx)
}

func (a *Users) Start() error {
	return a.grpcServer.Start()
}

func (a *Users) Stop() error {
	a.grpcServer.Stop()
	a.psql.Close()

	return nil
}

func (a *Users) MigrationUp(ctx context.Context, email, username, password string) error {
	return a.psql.MigrateUp(ctx, email, username, password)
}

func (a *Users) MigrationDown() error {
	return a.psql.MigrateDown()
}

func (a *Users) syncPermissions(ctx context.Context) error {
	roleRepo := a.psql.Role()
	permRepo := a.psql.Permissions()
	rolePermRepo := a.psql.RolePermissions()

	adminRole, err := roleRepo.GetRoleByName(ctx, "admin")
	if err != nil {
		return err
	}

	for _, perm := range a.Service.Permissions {
		logger.Info("syncing permission...", "scope", perm.Scope, "action", perm.Action)

		permission, err := permRepo.GetPermission(ctx, perm.Scope, perm.Action)
		if err != nil {
			permission, err = permRepo.AddPermission(ctx, perm.Name, perm.Description, perm.Scope, perm.Action)
			if err != nil {
				return err
			}
		}

		if err := rolePermRepo.AddRolePermissions(ctx, adminRole.ID, permission.ID); err != nil {
			return err
		}

		logger.Info("Bound permission", "permission_id", permission.ID, "role", adminRole.Name)
		logger.Info("synced permission", "scope", perm.Scope, "action", perm.Action)
	}

	return nil
}
