package main

import (
	"github.com/ezex-io/ezex-users/internal/utils/prompt"
	"github.com/spf13/cobra"
)

var configPath string

func main() {
	rootCmd := &cobra.Command{
		Use:               "ezex-users",
		Short:             "ezeX-users Service",
		CompletionOptions: cobra.CompletionOptions{HiddenDefaultCmd: true},
	}

	// Hide the "help" sub-command
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	rootCmd.PersistentFlags().StringVar(&configPath, "config", "",
		"Path to TOML config file (required)")
	_ = rootCmd.MarkPersistentFlagRequired("config")

	// register commands
	buildDBCMD(rootCmd)
	buildInit(rootCmd)
	buildStart(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		prompt.PrintErrorMsgf("%s", err)
	}
}
