// Package main provides the main entry point for the application.
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ezex-io/ezex-users/internal/config"
	"github.com/ezex-io/ezex-users/internal/controller"
	"github.com/ezex-io/ezex-users/internal/core/port/service"
	"github.com/ezex-io/ezex-users/internal/core/server"
	"github.com/ezex-io/ezex-users/internal/infra/repository"
	userspb "github.com/ezex-io/ezex-users/pkg/grpc"
	"github.com/ezex-io/gopkg/logger"
	"google.golang.org/grpc"
)

func main() {
	log := logger.NewSlog(nil)
	cfg, err := config.Load()
	if err != nil {
		log.Error("Failed to load config", "error", err)

		return
	}

	repo := repository.NewRepository()

	svc := service.NewService(repo)

	grpcServer := server.NewGRPCServer(cfg.GRPCServerAddress)

	log.Info("Starting gRPC server", "address", cfg.GRPCServerAddress)
	if err := grpcServer.Start(
		func(s *grpc.Server) {
			userspb.RegisterUsersServiceServer(s, controller.NewUserServer(svc.User()))
		},
	); err != nil {
		log.Error("Failed to start gRPC server", "error", err)

		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	go func() {
		grpcServer.Stop()
	}()

	<-ctx.Done()
}
