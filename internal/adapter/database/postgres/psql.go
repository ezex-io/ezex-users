package postgres

import (
	"context"
	"time"

	"github.com/ezex-io/ezex-users/internal/adapter/database/postgres/gen"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	db   *gen.Queries
	conn *pgxpool.Pool
}

func New(cfg *Config) (*Postgres, error) {
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

	return &Postgres{
		conn: conn,
		db:   gen.New(conn),
	}, nil
}

func (p *Postgres) Close() {
	p.conn.Close()
}

func (p *Postgres) Query() *gen.Queries {
	return p.db
}

func (p *Postgres) GetPool() *pgxpool.Pool {
	return p.conn
}
