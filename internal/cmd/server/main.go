// Package main provides the main entry point for the application.
package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/ezex-io/ezex-proto/go/users"
	"github.com/ezex-io/ezex-users/internal/adapter/database"
	"github.com/ezex-io/ezex-users/internal/adapter/database/postgres"
	"github.com/ezex-io/ezex-users/internal/adapter/grpc"
	"github.com/ezex-io/ezex-users/internal/interactor"
	"github.com/ezex-io/gopkg/env"
	"github.com/ezex-io/gopkg/logger"
	grp "google.golang.org/grpc"
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

	grpcServer := grpc.NewServer(cfg.GRPC, logging)

	secImageDb := database.NewSecurityImage(dbs.Query())
	userDB := database.NewUser(dbs.Query())

	secImageInteractor := interactor.NewSecurityImage(secImageDb)
	authInteractor := interactor.NewAuth(userDB)

	grpcServer.Register(func(s *grp.Server) {
		users.RegisterUsersServiceServer(s, grpc.NewUsersService(secImageInteractor, authInteractor))
		logging.Debug("user service registered")
	})

	logging.Info("Starting gRPC grpc", "address", cfg.GRPC.Address)
	if err := grpcServer.Start(); err != nil {
		logging.Fatal("Failed to start gRPC grpc", "error", err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-sig:
		logging.Warn("Received signal, shutting down...", "signal", s)
		grpcServer.Stop()
		dbs.Close()
	case err := <-grpcServer.Notify():
		logging.Error("gRPC grpc server error", "error", err)
		dbs.Close()
	}
}
