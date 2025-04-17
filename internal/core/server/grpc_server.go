// Package server provides server implementations for the application.
package server

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	server *grpc.Server
	addr   string
}

func NewGRPCServer(addr string) *GRPCServer {
	return &GRPCServer{
		server: grpc.NewServer(),
		addr:   addr,
	}
}

func (s *GRPCServer) Start(register func(*grpc.Server)) error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	register(s.server)

	if err := s.server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

func (s *GRPCServer) Stop() {
	s.server.GracefulStop()
}
