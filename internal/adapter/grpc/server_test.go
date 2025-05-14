package grpc_test

import (
	"testing"

	"github.com/ezex-io/ezex-proto/go/users"
	"github.com/ezex-io/ezex-users/internal/adapter/grpc"
	"github.com/ezex-io/gopkg/logger"
	"github.com/stretchr/testify/require"
	grp "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func TestServerStartupAndShutdown(t *testing.T) {
	logging := logger.NewSlog(nil)
	cfg := grpc.LoadFromEnv()
	require.NoError(t, cfg.BasicCheck())
	require.NotNil(t, cfg)

	cfg.EnableHealthCheck = true

	grpcServer := grpc.NewServer(cfg, logging)
	grpcServer.Register(func(s *grp.Server) {
		users.RegisterUsersServiceServer(s, grpc.NewUsersService(nil, nil))
	})

	go func() {
		err := grpcServer.Start()
		require.NoError(t, err)
	}()

	conn, err := grp.NewClient(cfg.Address, grp.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	client := grpc_health_v1.NewHealthClient(conn)
	res, err := client.Check(t.Context(), &grpc_health_v1.HealthCheckRequest{})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, res.Status, grpc_health_v1.HealthCheckResponse_SERVING)
	grpcServer.Stop()
}
