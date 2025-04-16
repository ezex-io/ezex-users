// Package main provides the main entry point for the application.
package main

import (
	"log"

	cmd "github.com/ezex-io/ezex-users/internal/cmd/server"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
