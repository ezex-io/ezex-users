package grpc

import (
	"context"

	"github.com/ezex-io/ezex-proto/go/users"
	"github.com/ezex-io/ezex-users/internal/entity"
)

func (s *UserServer) ProcessFirebaseLogin(ctx context.Context,
	req *users.ProcessFirebaseLoginRequest,
) (*users.ProcessFirebaseLoginResponse, error) {
	resp, err := s.auth.ProcessFirebaseLogin(ctx, &entity.ProcessFirebaseLoginRequest{
		Email:       req.Email,
		FirebaseUID: req.FirebaseUserId,
	})
	if err != nil {
		return nil, err
	}

	return &users.ProcessFirebaseLoginResponse{
		UserId: resp.UserID,
	}, nil
}
