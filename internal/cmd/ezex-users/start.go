package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ezex-io/ezex-users/internal/apps"
	"github.com/ezex-io/ezex-users/internal/config"
	"github.com/ezex-io/ezex-users/pkg/logger"
	"github.com/spf13/cobra"
)

func buildStart(root *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start artogenia api interactors",
	}

	root.AddCommand(cmd)

	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		cfg, err := config.Load(configPath)
		if err != nil {
			return err
		}

		users, err := apps.NewUsers(cmd.Context(), cfg, false)
		if err != nil {
			return err
		}

		if err := users.Start(); err != nil {
			return err
		}

		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

		s := <-interrupt

		_ = users.Stop()
		logger.Warn("app/run: signal received", "signal", s.String())

		return nil
	}

	return cmd
}
