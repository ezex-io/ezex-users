// Package main provides the main entry point for the application.
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ezex-io/ezex-users/internal/config"
	"github.com/ezex-io/ezex-users/internal/core/server"
	"github.com/ezex-io/ezex-users/internal/core/service"
	"github.com/ezex-io/ezex-users/internal/infra/repository"
	"github.com/ezex-io/gopkg/logger"
)

var log = logger.NewSlog(nil)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config", "error", err)
	}

	repo := repository.NewSecurityImageRepository()

	securityImageService := service.NewSecurityImageService(repo)

	grpcServer := server.NewGRPCServer(cfg.GRPCServerAddress, securityImageService)

	log.Info("Starting gRPC server", "address", cfg.GRPCServerAddress)
	if err := grpcServer.Start(); err != nil {
		log.Error("Failed to start gRPC server", "error", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := grpcServer.Stop(ctx); err != nil {
		log.Error("Failed to stop gRPC server", "error", err)
	}
}
