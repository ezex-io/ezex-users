package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/ezex-io/ezex-users/internal/core/common/router"
	"github.com/ezex-io/ezex-users/internal/core/server"
	"github.com/ezex-io/ezex-users/internal/core/service"
	"github.com/ezex-io/ezex-users/internal/infra/config"
	"github.com/ezex-io/ezex-users/internal/infra/logger"
	"github.com/ezex-io/ezex-users/internal/infra/repository"
	"github.com/joho/godotenv"
)

func main() {
	envFile := flag.String("env", ".env.example", "Path to environment file")
	flag.Parse()

	logging := logger.NewSlog(nil)

	if err := godotenv.Load(*envFile); err != nil {
		logging.Warn("Failed to load env file '%s': %v. Continuing with system environment...", *envFile, err)
	} else {
		logging.Debug("Loaded environment variables from '%s'", *envFile)
	}

	cfg, err := config.Load()
	if err != nil {
		logging.Fatal("Failed to load config: %v", err)
	}

	if err := cfg.BasicCheck(); err != nil {
		logging.Fatal("Configuration validation failed: %v", err)
	}

	httpRouter := router.SetupRouter()
	httpServer := server.NewHTTPServer(cfg.HTTPServer.Address, httpRouter)

	securityImageService := service.NewSecurityImageService(
		repository.NewSecurityImageRepository(),
	)
	grpcServer := server.NewGRPCServer(cfg.GRPCServer.Address, securityImageService)

	go func() {
		logging.Info("Starting HTTP server on %s", cfg.HTTPServer.Address)
		if err := httpServer.Start(); err != nil {
			logging.Fatal("Failed to start HTTP server: %v", err)
		}
	}()

	go func() {
		logging.Info("Starting gRPC server on %s", cfg.GRPCServer.Address)
		if err := grpcServer.Start(); err != nil {
			logging.Fatal("Failed to start gRPC server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logging.Info("Shutting down servers...")

	ctx, cancel := context.WithTimeout(context.Background(), 5)
	defer cancel()
	if err := httpServer.Stop(ctx); err != nil {
		logging.Fatal("Failed to stop HTTP server: %v", err)
	}

	if err := grpcServer.Stop(ctx); err != nil {
		logging.Fatal("Failed to stop gRPC server: %v", err)
	}

	logging.Info("Servers stopped")
}
