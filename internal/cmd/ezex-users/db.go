package main

import (
	"fmt"

	"github.com/ezex-io/ezex-users/internal/adapters/db/postgres"
	"github.com/ezex-io/ezex-users/internal/config"
	"github.com/ezex-io/ezex-users/internal/utils/prompt"
	"github.com/spf13/cobra"
)

func buildDBCMD(root *cobra.Command) *cobra.Command {
	dbCMD := &cobra.Command{
		Use:   "db",
		Short: "Database management commands",
		Long:  "Run and manage database migrations and initialization tasks.",
	}

	dbCMD.AddCommand(buildMigrationCMD())
	root.AddCommand(dbCMD)

	return dbCMD
}

func buildMigrationCMD() *cobra.Command {
	mgCMD := &cobra.Command{
		Use:     "migration",
		Aliases: []string{"migrate"},
		Short:   "Run database migrations",
		Long:    "Apply or revert SQL migrations defined in the migrations/ directory.",
	}

	mgCMD.AddCommand(migrationUpCmd())
	mgCMD.AddCommand(migrationDownCmd())

	return mgCMD
}

func migrationUpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up",
		Short: "Apply all up migrations",
	}

	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		cfg, err := config.Load(configPath)
		if err != nil {
			return fmt.Errorf("load config: %w", err)
		}
		if err := cfg.BasicCheck(); err != nil {
			return fmt.Errorf("invalid config: %w", err)
		}

		psql, err := postgres.New(cmd.Context(), cfg.Database)
		if err != nil {
			return fmt.Errorf("connect database: %w", err)
		}
		defer psql.Close()

		email := prompt.Input("email")
		username := prompt.Input("username")
		password := prompt.Input("password")

		return psql.MigrateUp(cmd.Context(), email, username, password)
	}

	return cmd
}

func migrationDownCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "down",
		Short: "Revert all migrations",
		RunE: func(cmd *cobra.Command, _ []string) error {
			cfg, err := config.Load(configPath)
			if err != nil {
				return fmt.Errorf("load config: %w", err)
			}
			if err := cfg.BasicCheck(); err != nil {
				return fmt.Errorf("invalid config: %w", err)
			}

			psql, err := postgres.New(cmd.Context(), cfg.Database)
			if err != nil {
				return fmt.Errorf("connect database: %w", err)
			}
			defer psql.Close()

			return psql.MigrateDown()
		},
	}
}
