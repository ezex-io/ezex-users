package grpc

import (
	"fmt"
	"net"

	"github.com/Sensifai-BV/artogenia/pkg/logger"
	"github.com/ezex-io/ezex-proto/go/users"
	"github.com/ezex-io/ezex-users/internal/interactors/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	users.UnimplementedUsersServiceServer

	srv     *grpc.Server
	cfg     *Config
	logging *logger.SubLogger

	userInteractor *user.User
}

func New(cfg *Config, userInteractor *user.User) *Server {
	server := &Server{
		srv:            grpc.NewServer(),
		cfg:            cfg,
		logging:        logger.NewSubLogger("_grpc", nil),
		userInteractor: userInteractor,
	}

	users.RegisterUsersServiceServer(server.srv, server)

	if cfg.EnableReflection {
		reflection.Register(server.srv)
		server.logging.Debug("Reflection enabled")
	}

	if cfg.EnableHealth {
		grpc_health_v1.RegisterHealthServer(server.srv, health.NewServer())
		server.logging.Debug("gRPC health enabled")
	}

	return server
}

func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", s.cfg.Address, s.cfg.Port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	go func() {
		if err := s.srv.Serve(lis); err != nil {
			s.logging.Fatal(err.Error())
		}
	}()

	s.logging.Info("gRPC server started", "addr", addr)

	return nil
}

func (s *Server) Stop() {
	s.srv.GracefulStop()
}
