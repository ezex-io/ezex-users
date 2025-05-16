// Package main provides the main entry point for the application.
package main

import (
	"context"
	"flag"

	"github.com/ezex-io/ezex-users/internal/adapter/database"
	"github.com/ezex-io/ezex-users/internal/adapter/database/postgres"
	"github.com/ezex-io/ezex-users/internal/adapter/grpc"
	"github.com/ezex-io/ezex-users/internal/interactor"
	"github.com/ezex-io/gopkg/env"
	"github.com/ezex-io/gopkg/logger"
	"github.com/ezex-io/gopkg/utils"
)

func main() {
	envFile := flag.String("env", ".env", "Path to environment file")
	flag.Parse()

	logging := logger.NewSlog(nil)

	if err := env.LoadEnvsFromFile(*envFile); err != nil {
		logging.Debug("Failed to load env file '%s': %v. Continuing with system environment...", *envFile, err)
	}

	cfg := makeConfig()
	if err := cfg.BasicCheck(); err != nil {
		logging.Fatal("failed to load config: %v", err)
	}

	dbs, err := postgres.New(cfg.Database)
	if err != nil {
		logging.Fatal("failed to connect to database", "err", err)
	}

	// run migrations
	if err := postgres.MigrateDB(context.Background(), dbs.GetPool()); err != nil {
		logging.Fatal("failed to run migrations", "err", err)
	}
	logging.Info("Database migrations completed successfully")

	secImageDb := database.NewSecurityImage(dbs.Query())
	userDB := database.NewUser(dbs.Query())

	secImageInteractor := interactor.NewSecurityImage(secImageDb)
	authInteractor := interactor.NewAuth(userDB)

	usersService := grpc.NewUsersService(secImageInteractor, authInteractor)

	server, err := grpc.NewServer(cfg.GRPC, logging, usersService)
	if err != nil {
		logging.Fatal("failed to create gRPC server: %v", err)
	}

	server.Start()

	utils.TrapSignal(func() {
		logging.Info("Exiting...")

		server.Stop()
	})

	// run forever
	select {}
}
