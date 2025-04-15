package main

import (
	"context"
	"testing"
	"time"

	"github.com/ezex-io/ezex-users/internal/core/common/router"
	"github.com/ezex-io/ezex-users/internal/core/server"
	"github.com/ezex-io/ezex-users/internal/core/service"
	"github.com/ezex-io/ezex-users/internal/infra/config"
	"github.com/ezex-io/ezex-users/internal/infra/repository"
)

func TestServerStartupAndShutdown(t *testing.T) {
	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	httpRouter := router.SetupRouter()
	httpServer := server.NewHTTPServer(cfg.HTTPServerAddress, httpRouter)

	securityImageService := service.NewSecurityImageService(
		repository.NewSecurityImageRepository(),
	)
	grpcServer := server.NewGRPCServer(cfg.GRPCServerAddress, securityImageService)

	httpErr := make(chan error, 1)
	grpcErr := make(chan error, 1)

	go func() {
		httpErr <- httpServer.Start()
	}()

	go func() {
		grpcErr <- grpcServer.Start()
	}()

	time.Sleep(100 * time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Stop(ctx); err != nil {
		t.Errorf("Failed to stop HTTP server: %v", err)
	}

	if err := grpcServer.Stop(ctx); err != nil {
		t.Errorf("Failed to stop gRPC server: %v", err)
	}

	select {
	case err := <-httpErr:
		if err != nil && err.Error() != "http: Server closed" {
			t.Errorf("HTTP server error: %v", err)
		}
	default:
	}

	select {
	case err := <-grpcErr:
		if err != nil && err.Error() != "grpc: the server has been stopped" {
			t.Errorf("gRPC server error: %v", err)
		}
	default:
	}
}
