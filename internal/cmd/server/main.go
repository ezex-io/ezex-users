// Package cmd provides the main entry point for the application.
package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ezex-io/ezex-users/internal/config"
	"github.com/ezex-io/ezex-users/internal/core/port/service"
	"github.com/ezex-io/ezex-users/internal/core/server"
	"github.com/ezex-io/ezex-users/internal/infra/repository"
	"github.com/ezex-io/gopkg/logger"
)

var log = logger.NewSlog(nil)

func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	repo := repository.NewRepository()

	svc := service.NewService(repo)

	grpcServer := server.NewGRPCServer(cfg.GRPCServerAddress, svc)

	log.Info("Starting gRPC server", "address", cfg.GRPCServerAddress)
	if err := grpcServer.Start(); err != nil {
		log.Error("Failed to start gRPC server", "error", err)

		return fmt.Errorf("failed to start gRPC server: %w", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := grpcServer.Stop(ctx); err != nil {
		log.Error("Failed to stop gRPC server", "error", err)

		return fmt.Errorf("failed to stop gRPC server: %w", err)
	}

	return nil
}
