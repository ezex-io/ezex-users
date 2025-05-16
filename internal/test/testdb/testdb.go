package testdb

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/ezex-io/ezex-users/internal/adapter/database/postgres"
	"github.com/ezex-io/ezex-users/internal/adapter/database/postgres/gen"
	"github.com/ezex-io/gopkg/env"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

var (
	setupOnce sync.Once
	envLoaded bool
)

func Setup() {
	setupOnce.Do(func() {
		// load from project root
		envFile := "../../../.env.example"
		if err := env.LoadEnvsFromFile(envFile); err != nil {
			log.Printf("Failed to load %s: %v. Using default test settings.\n", envFile, err)
		} else {
			log.Printf("Successfully loaded %s\n", envFile)
			envLoaded = true
		}
	})

	if envLoaded {
		log.Println("Environment already loaded")
	}
}

type PostgresDB struct {
	DB      *postgres.Postgres
	Queries *gen.Queries
	Pool    *pgxpool.Pool
	t       *testing.T
}

func getTestDBConfig(t *testing.T) *postgres.Config {
	t.Helper()

	envConfig := postgres.LoadFromEnv()
	defaultHost := "localhost:5433"
	username := "postgres"
	password := "postgres"
	database := "postgres"

	if envConfig == nil {
		t.Log("Warning: No environment configuration found, using defaults")

		return &postgres.Config{
			Address:  defaultHost,
			Username: username,
			Password: password,
			Database: database,
		}
	}

	if envConfig.Address == "" {
		envConfig.Address = defaultHost
	}

	if envConfig.Username == "" {
		envConfig.Username = username
	}

	if envConfig.Password == "" {
		envConfig.Password = password
	}

	if envConfig.Database == "" {
		envConfig.Database = database
	}

	return envConfig
}

func NewPostgresDB(t *testing.T) *PostgresDB {
	t.Helper()

	Setup()

	dbName := fmt.Sprintf("test_db_%s", uuid.New().String()[:8])

	envConfig := getTestDBConfig(t)

	dbURL := buildConnectionString(envConfig.Address, "postgres", envConfig.Username, envConfig.Password)

	ctx, cancel := context.WithTimeout(t.Context(), time.Second*30)
	defer cancel()

	conn, err := pgxpool.New(ctx, dbURL)
	require.NoError(t, err)

	_, err = conn.Exec(ctx, fmt.Sprintf("CREATE DATABASE %s", dbName))
	require.NoError(t, err)

	conn.Close()

	// Create a config for the test database using the same credentials
	cfg := &postgres.Config{
		Address:         envConfig.Address,
		Database:        dbName, // Use the newly created test database
		Username:        envConfig.Username,
		Password:        envConfig.Password,
		MaxOpenConns:    5,
		MaxIdleConns:    2,
		ConnMaxLifetime: "5m",
	}

	dbConn, err := postgres.New(cfg)
	require.NoError(t, err)

	err = postgres.MigrateDB(ctx, dbConn.GetPool())
	require.NoError(t, err)

	return &PostgresDB{
		DB:      dbConn,
		Queries: dbConn.Query(),
		Pool:    dbConn.GetPool(),
		t:       t,
	}
}

func (tdb *PostgresDB) Cleanup() {
	tdb.t.Helper()

	pool := tdb.Pool
	dbName := pool.Config().ConnConfig.Database

	tdb.DB.Close()

	config := getTestDBConfig(tdb.t)

	dbURL := buildConnectionString(config.Address, "postgres", config.Username, config.Password)

	ctx, cancel := context.WithTimeout(tdb.t.Context(), time.Second*5)
	defer cancel()

	conn, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		tdb.t.Logf("Failed to connect to postgres for cleanup: %v", err)

		return
	}
	defer conn.Close()

	_, err = conn.Exec(ctx, fmt.Sprintf("DROP DATABASE %s WITH (FORCE)", dbName))
	if err != nil {
		tdb.t.Logf("Failed to drop test database: %v", err)
	}
}

func buildConnectionString(address, database, username, password string) string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s", username, password, address, database)
}
