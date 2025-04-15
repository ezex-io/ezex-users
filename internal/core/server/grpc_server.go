package server

import (
	"context"
	"fmt"
	"net"

	security_imagev1 "github.com/ezex-io/ezex-users/api/gen/go/security_image/v1"
	"github.com/ezex-io/ezex-users/internal/core/port/service"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	server  *grpc.Server
	address string
}

func NewGRPCServer(address string, securityImageService service.SecurityImageService) *GRPCServer {
	s := grpc.NewServer()
	security_imagev1.RegisterSecurityImageServiceServer(s, NewSecurityImageServer(securityImageService))

	return &GRPCServer{
		server:  s,
		address: address,
	}
}

func (s *GRPCServer) Start() error {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	return s.server.Serve(lis)
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
		return ctx.Err()
	case <-done:
		return nil
	}
}
