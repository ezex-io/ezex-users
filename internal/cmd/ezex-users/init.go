package main

import (
	"github.com/ezex-io/ezex-users/internal/apps"
	"github.com/ezex-io/ezex-users/internal/config"
	"github.com/ezex-io/ezex-users/internal/utils/prompt"
	"github.com/spf13/cobra"
)

func buildInit(root *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize ezeX-users interactors from scratch",
	}

	root.AddCommand(cmd)

	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		confirmed := prompt.Confirm("Do you wish to initialize ezeX-users interactors?")
		if !confirmed {
			return nil
		}

		prompt.PrintLine()

		email := prompt.Input("email")
		username := prompt.Input("username")
		password := prompt.Password("password", true)

		cfg, err := config.Load(configPath)
		if err != nil {
			return err
		}

		users, err := apps.NewUsers(cmd.Context(), cfg, true)
		if err != nil {
			return err
		}

		if err := users.MigrationUp(cmd.Context(), email, username, password); err != nil {
			return err
		}

		return users.Init(cmd.Context())
	}

	return cmd
}
