package grpc

import (
	"fmt"
	"net"

	"github.com/ezex-io/ezex-proto/go/users"
	"github.com/ezex-io/gopkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
	logger   logger.Logger
}

func NewServer(conf *Config, logger logger.Logger, usersService *UsersService) (*Server, error) {
	listener, err := net.Listen("tcp", conf.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on port: %w", err)
	}

	server := grpc.NewServer()
	users.RegisterUsersServiceServer(server, usersService)

	if conf.EnableHealthCheck {
		grpc_health_v1.RegisterHealthServer(server, health.NewServer())
		logger.Info("gRPC health check enabled")
	}

	if conf.EnableReflection {
		reflection.Register(server)
		logger.Info("gRPC reflection enabled")
	}

	return &Server{
		server:   server,
		listener: listener,
		logger:   logger,
	}, nil
}

func (s *Server) Start() {
	go func() {
		s.logger.Info("gRPC server start listening", "address", s.listener.Addr())
		if err := s.server.Serve(s.listener); err != nil {
			s.logger.Debug("error on gRPC server", "error", err)
		}
	}()
}

func (s *Server) Stop() {
	if s.server != nil {
		s.server.Stop()
		_ = s.listener.Close()
	}
}
