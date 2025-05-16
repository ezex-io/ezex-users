package postgres

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Embeds all SQL migration files
//
//go:embed migrations/*.sql
var migrationFS embed.FS

// MigrateDB applies all migrations to the database.
func MigrateDB(ctx context.Context, conn *pgxpool.Pool) error {
	// create migrations table if it doesn't exist
	_, err := conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version TEXT PRIMARY KEY,
			applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			err_description TEXT
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	// get applied migrations
	rows, err := conn.Query(ctx, "SELECT version, err_description FROM schema_migrations")
	if err != nil {
		return fmt.Errorf("failed to get applied migrations: %w", err)
	}
	defer rows.Close()

	appliedMigrations := make(map[string]bool)
	failedMigrations := make(map[string]string)
	for rows.Next() {
		var version string
		var errDescription *string
		if err := rows.Scan(&version, &errDescription); err != nil {
			return fmt.Errorf("failed to scan migration version: %w", err)
		}
		appliedMigrations[version] = true

		if errDescription != nil {
			failedMigrations[version] = *errDescription
			log.Printf("Previously failed migration: %s - Error: %s", version, *errDescription)
		}
	}

	// read all migration files
	entries, err := fs.ReadDir(migrationFS, "migrations")
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %w", err)
	}

	upMigrations := make(map[string]string)
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".sql") {
			continue
		}

		content, err := fs.ReadFile(migrationFS, "migrations/"+entry.Name())
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", entry.Name(), err)
		}

		key := strings.TrimSuffix(entry.Name(), ".sql")
		if strings.HasSuffix(key, ".up") {
			baseName := strings.TrimSuffix(key, ".up")
			upMigrations[baseName] = string(content)
		} else if !strings.HasSuffix(key, ".down") {
			// handle migrations without up/down suffix (legacy format)
			upMigrations[key] = string(content)
		}
	}

	// sort migrations
	orderedMigrations := sortMigrations(keysFromMap(upMigrations))

	// apply migrations in a transaction
	transaction, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err := transaction.Rollback(ctx); err != nil {
			log.Printf("failed to rollback transaction: %s", err)
		}
	}()

	for _, name := range orderedMigrations {
		// skip successful migrations
		if appliedMigrations[name] && failedMigrations[name] == "" {
			continue
		}

		// apply migration
		sql := upMigrations[name]
		_, err = transaction.Exec(ctx, sql)
		if err != nil {
			// record failure and return
			_ = updateMigrationStatus(ctx, transaction, name, err.Error())

			return fmt.Errorf("failed to apply migration %s: %w", name, err)
		}

		// record success
		_ = updateMigrationStatus(ctx, transaction, name, "")

		log.Printf("Applied migration: %s", name)
	}

	return transaction.Commit(ctx)
}

// helper function to record migration status.
func updateMigrationStatus(ctx context.Context, transaction pgx.Tx, version, errDescription string) error {
	query := `
		INSERT INTO schema_migrations (version, err_description) 
		VALUES ($1, $2) 
		ON CONFLICT (version) DO UPDATE 
		SET err_description = $2, applied_at = NOW()
	`

	_, err := transaction.Exec(ctx, query, version, stringToNilablePtr(errDescription))

	return err
}

// convert empty string to nil, otherwise return pointer to string.
func stringToNilablePtr(s string) *string {
	if s == "" {
		return nil
	}

	return &s
}

// get keys from a map.
func keysFromMap(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// sort migrations.
func sortMigrations(migrations []string) []string {
	// simple sorting
	for i := 0; i < len(migrations); i++ {
		for j := i + 1; j < len(migrations); j++ {
			if migrations[i] > migrations[j] {
				migrations[i], migrations[j] = migrations[j], migrations[i]
			}
		}
	}

	return migrations
}

// ResetDB clears all data from tables (for integration tests ONLY).
func ResetDB(ctx context.Context, conn *pgxpool.Pool) error {
	_, err := conn.Exec(ctx, "TRUNCATE users CASCADE")

	return err
}
