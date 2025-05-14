package grpc

import (
	"fmt"
	"net"

	"github.com/ezex-io/gopkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	server  *grpc.Server
	logging logger.Logger
	addr    string
	errCh   chan error
}

func NewServer(cfg *Config, logging logger.Logger) *Server {
	srv := &Server{
		server:  grpc.NewServer(),
		addr:    cfg.Address,
		errCh:   make(chan error),
		logging: logging,
	}

	if cfg.EnableHealthCheck {
		grpc_health_v1.RegisterHealthServer(srv.server, health.NewServer())
		logging.Info("gRPC health check enabled")
	}

	if cfg.EnableReflection {
		reflection.Register(srv.server)
		logging.Info("gRPC reflection enabled")
	}

	return srv
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	go func() {
		s.errCh <- s.server.Serve(listener)
	}()

	return nil
}

func (s *Server) Register(regFunc func(s *grpc.Server)) {
	regFunc(s.server)
}

func (s *Server) Notify() <-chan error {
	return s.errCh
}

func (s *Server) Stop() {
	s.server.GracefulStop()
}
