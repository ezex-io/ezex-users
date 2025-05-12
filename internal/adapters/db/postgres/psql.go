package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ezex-io/ezex-users/internal/adapters/db/postgres/gen"
	"github.com/ezex-io/ezex-users/internal/ports"
	"github.com/ezex-io/ezex-users/internal/ports/repository"
	"github.com/ezex-io/ezex-users/internal/utils/bcrypt"
	"github.com/ezex-io/ezex-users/migrations"
	"github.com/ezex-io/ezex-users/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/segmentio/ksuid"
)

var _ ports.PostgresPort = (*Postgres)(nil)

type Postgres struct {
	cfg  *Config
	conn *pgxpool.Pool
	db   gen.Querier

	logging *logger.SubLogger

	userRepo            repository.UserPort
	roleRepo            repository.RolePort
	permissionRepo      repository.PermissionsPort
	rolePermissionsRepo repository.RolePermissionsPort
	securityImageRepo   repository.SecurityImagePort
}

func New(ctx context.Context, cfg *Config) (ports.PostgresPort, error) {
	connString := cfg.URI()

	poolCfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	poolCfg.MaxConns = cfg.MaxOpenConns
	poolCfg.MinConns = cfg.MaxIdleConns

	if d, err := time.ParseDuration(cfg.ConnMaxLifetime); err == nil {
		poolCfg.MaxConnLifetime = d
	}

	conn, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, err
	}

	pingCtx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	if err := conn.Ping(pingCtx); err != nil {
		return nil, err
	}

	psql := &Postgres{
		cfg:  cfg,
		conn: conn,
		db:   gen.New(conn),
	}

	psql.logging = logger.NewSubLogger("_database", psql)

	psql.init()

	return psql, nil
}

func (p *Postgres) init() {
	p.userRepo = newUserRepo(p.db)
	p.roleRepo = newRoleRepo(p.db)
	p.permissionRepo = newPermissionsRepo(p.db)
	p.rolePermissionsRepo = newRolePermissionsRepo(p.db)
	p.securityImageRepo = newSecurityImageRepo()
}

func (p *Postgres) Close() {
	p.conn.Close()
}

func (*Postgres) String() string {
	return "Postgres"
}

func (p *Postgres) User() repository.UserPort {
	return p.userRepo
}

func (p *Postgres) Role() repository.RolePort {
	return p.roleRepo
}

func (p *Postgres) Permissions() repository.PermissionsPort {
	return p.permissionRepo
}

func (p *Postgres) RolePermissions() repository.RolePermissionsPort {
	return p.rolePermissionsRepo
}

func (p *Postgres) SecurityImage() repository.SecurityImagePort {
	return p.securityImageRepo
}

func (p *Postgres) MigrateUp(ctx context.Context, email, username, password string) error {
	if email == "" {
		return errors.New("email is required for migration")
	}

	if username == "" {
		return errors.New("username is required for migration")
	}

	if password == "" {
		return errors.New("password is required for migration")
	}

	mig, err := p.migrate()
	if err != nil {
		return err
	}
	defer func() {
		_, _ = mig.Close()
	}()

	p.logging.Info("prepare schema version...")

	ver, dirty, err := mig.Version()
	if !errors.Is(err, migrate.ErrNilVersion) && err != nil {
		return err
	}

	if ver != 0 {
		p.logging.Info("found database version", "version", ver)
	}

	if dirty {
		if err = mig.Migrate(ver); err != nil {
			return fmt.Errorf("error migrating to version %d resetting dirty: %w", ver, err)
		}
	}

	p.logging.Info("starting migrations...")

	if err := mig.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("running migrations: %w", err)
	}

	userID := ksuid.New().String()
	adminRoleID := ksuid.New().String()

	roles := []gen.CreateRolesParams{
		{
			ID:        adminRoleID,
			Name:      "admin",
			IsSystem:  true,
			IsDefault: false,
		},
		{
			ID:        ksuid.New().String(),
			Name:      "user",
			IsSystem:  true,
			IsDefault: true,
		},
		{
			ID:        ksuid.New().String(),
			Name:      "support",
			IsSystem:  true,
			IsDefault: false,
		},
	}

	p.logging.Info("creating default roles...")

	res, err := p.db.CreateRoles(ctx, roles)
	if err != nil {
		return err
	}

	if int(res) < len(roles) {
		return fmt.Errorf("expected to create %d roles, got %d", len(roles), res)
	}

	p.logging.Info("roles admin, user, support successfully created.")

	hashedPass := bcrypt.MustHash(password)

	p.logging.Info("creating default user...", "email", email, "username", username)

	if err := p.db.CreateUserWithID(ctx, gen.CreateUserWithIDParams{
		ID:          userID,
		Username:    username,
		Email:       email,
		Password:    pgtype.Text{String: hashedPass, Valid: true},
		Status:      gen.UsersStatusEnumACTIVE,
		CreatedByID: pgtype.Text{String: userID, Valid: true},
	}); err != nil {
		return err
	}

	p.logging.Info("default user successfully created.")
	p.logging.Info("granting role to user...")

	if err := p.db.GrantRole(ctx, gen.GrantRoleParams{
		UserID:      userID,
		RoleID:      adminRoleID,
		GrantedByID: userID,
		Email:       email,
	}); err != nil {
		return err
	}

	p.logging.Info("migration successfully done.")

	return nil
}

func (p *Postgres) MigrateDown() error {
	mig, err := p.migrate()
	if err != nil {
		return err
	}
	defer func() {
		_, _ = mig.Close()
	}()

	p.logging.Info("prepare schema version...")

	version, dirty, err := mig.Version()
	if errors.Is(err, migrate.ErrNilVersion) {
		p.logging.Debug("No migrations have been applied (version 0). Nothing to revert.")

		return nil
	}
	if dirty {
		return fmt.Errorf("database is in dirty state at version %d", version)
	}

	p.logging.Info("starting migrations...")

	return mig.Down()
}

func (p *Postgres) migrate() (*migrate.Migrate, error) {
	p.logging.Info("initializing migrations...")

	dbs, err := sql.Open("pgx", p.cfg.URI())
	if err != nil {
		return nil, err
	}

	dbs.SetMaxIdleConns(int(p.cfg.MaxIdleConns))
	dbs.SetMaxOpenConns(int(p.cfg.MaxOpenConns))

	p.logging.Info("creating database driver...")

	driver, err := pgx.WithInstance(dbs, &pgx.Config{})
	if err != nil {
		return nil, fmt.Errorf("creating PGX driver: %w", err)
	}

	p.logging.Info("prepare schema files...")

	src, err := iofs.New(migrations.PSQLEmbed, "psql")
	if err != nil {
		return nil, fmt.Errorf("creating I/O FS source: %w", err)
	}

	mig, err := migrate.NewWithInstance("iofs", src, "postgres", driver)
	if err != nil {
		return nil, fmt.Errorf("creating migrate instance: %w", err)
	}

	p.logging.Info("creating migrations...")

	return mig, nil
}
