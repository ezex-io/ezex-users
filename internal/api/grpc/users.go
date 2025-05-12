package grpc

import (
	"context"

	"github.com/ezex-io/ezex-proto/go/users"
)

func (s *Server) SaveSecurityImage(ctx context.Context,
	in *users.SaveSecurityImageRequest,
) (*users.SaveSecurityImageResponse, error) {
	return s.userInteractor.SaveSecurityImage(ctx, in)
}

func (s *Server) GetSecurityImage(ctx context.Context,
	in *users.GetSecurityImageRequest,
) (*users.GetSecurityImageResponse, error) {
	return s.userInteractor.GetSecurityImage(ctx, in)
}
