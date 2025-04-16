// Package server provides server implementations for HTTP and gRPC.
package server

import (
	"context"
	"fmt"
	"net"

	"github.com/ezex-io/ezex-users/api/grpc/proto"
	"github.com/ezex-io/ezex-users/internal/core/port/service"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	server  *grpc.Server
	address string
}

func NewGRPCServer(address string, userService service.UserService) *GRPCServer {
	s := grpc.NewServer()
	proto.RegisterUsersServiceServer(s, NewUserServer(userService))

	return &GRPCServer{
		server:  s,
		address: address,
	}
}

func (s *GRPCServer) Start() error {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	if err := s.server.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

func (s *GRPCServer) Stop(ctx context.Context) error {
	done := make(chan bool)
	go func() {
		s.server.GracefulStop()
		done <- true
	}()

	select {
	case <-ctx.Done():
		s.server.Stop()

		return fmt.Errorf("failed to stop gRPC server: %w", ctx.Err())
	case <-done:
		return nil
	}
}
