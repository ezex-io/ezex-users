package cmd

import (
	"context"
	"testing"
	"time"

	"github.com/ezex-io/ezex-users/internal/config"
	"github.com/ezex-io/ezex-users/internal/core/port/service"
	"github.com/ezex-io/ezex-users/internal/core/server"
	"github.com/ezex-io/ezex-users/internal/infra/repository"
)

func TestServerStartupAndShutdown(t *testing.T) {
	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	service := service.NewService(repository.NewRepository())

	grpcServer := server.NewGRPCServer(cfg.GRPCServerAddress, service)

	grpcErr := make(chan error, 1)

	go func() {
		grpcErr <- grpcServer.Start()
	}()

	time.Sleep(100 * time.Millisecond)

	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()

	if err := grpcServer.Stop(ctx); err != nil {
		t.Errorf("Failed to stop gRPC server: %v", err)
	}

	select {
	case err := <-grpcErr:
		if err != nil && err.Error() != "grpc: the server has been stopped" {
			t.Errorf("gRPC server error: %v", err)
		}
	default:
	}
}
