package grpc

import (
	"context"

	"github.com/ezex-io/ezex-proto/go/users"
	"github.com/ezex-io/ezex-users/internal/port"
)

func (s *UserServer) ProcessLogin(ctx context.Context,
	req *users.ProcessLoginRequest,
) (*users.ProcessLoginResponse, error) {
	res, err := s.auth.ProcessLogin(ctx, &port.ProcessLoginRequest{
		Email:       req.Email,
		FirebaseUID: req.FirebaseUserId,
	})
	if err != nil {
		return nil, err
	}

	return &users.ProcessLoginResponse{
		UserId: res.UserID,
	}, nil
}
