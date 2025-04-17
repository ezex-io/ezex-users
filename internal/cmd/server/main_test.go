package main

import (
	"testing"
	"time"

	"github.com/ezex-io/ezex-users/api/grpc/proto"
	"github.com/ezex-io/ezex-users/internal/config"
	"github.com/ezex-io/ezex-users/internal/controller"
	"github.com/ezex-io/ezex-users/internal/core/port/service"
	"github.com/ezex-io/ezex-users/internal/core/server"
	"github.com/ezex-io/ezex-users/internal/infra/repository"
	"google.golang.org/grpc"
)

func TestServerStartupAndShutdown(t *testing.T) {
	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	service := service.NewService(repository.NewRepository())

	grpcServer := server.NewGRPCServer(cfg.GRPCServerAddress)

	grpcErr := make(chan error, 1)

	go func() {
		grpcErr <- grpcServer.Start(func(s *grpc.Server) {
			proto.RegisterUserServiceServer(s, controller.NewUserServer(service.User()))
		})
	}()

	time.Sleep(100 * time.Millisecond)

	grpcServer.Stop()

	select {
	case err := <-grpcErr:
		if err != nil && err.Error() != "grpc: the server has been stopped" {
			t.Errorf("gRPC server error: %v", err)
		}
	default:
	}
}
