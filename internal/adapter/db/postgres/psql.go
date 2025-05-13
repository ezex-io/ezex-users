package postgres

import (
	"context"
	"time"

	"github.com/ezex-io/ezex-users/internal/adapter/db/postgres/gen"
	"github.com/ezex-io/ezex-users/internal/port"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	db   *gen.Queries
	conn *pgxpool.Pool
}

func New(cfg *Config) (port.PostgresPort, error) {
	connString := cfg.uri()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

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

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return &postgres{
		conn: conn,
		db:   gen.New(conn),
	}, nil
}

func (p *postgres) Close() {
	p.conn.Close()
}

func (p *postgres) Query() *gen.Queries {
	return p.db
}
