package grpc

import (
	"context"
	"fmt"

	"github.com/ezex-io/ezex-proto/go/users"
	"github.com/ezex-io/ezex-users/internal/port"
)

func (s *UserServer) SaveSecurityImage(
	ctx context.Context,
	req *users.SaveSecurityImageRequest,
) (*users.SaveSecurityImageResponse, error) {
	_, err := s.securityImage.SaveSecurityImage(ctx, &port.SaveSecurityImageRequest{
		Email:          req.Email,
		SecurityImage:  req.SecurityImage,
		SecurityPhrase: req.SecurityPhrase,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to save security image: %w", err)
	}

	return &users.SaveSecurityImageResponse{}, nil
}

func (s *UserServer) GetSecurityImage(
	ctx context.Context,
	req *users.GetSecurityImageRequest,
) (*users.GetSecurityImageResponse, error) {
	res, err := s.securityImage.GetSecurityImage(ctx, &port.GetSecurityImageRequest{
		Email: req.Email,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get security image: %w", err)
	}

	return &users.GetSecurityImageResponse{
		SecurityImage:  res.SecurityImage,
		SecurityPhrase: res.SecurityPhrase,
	}, nil
}
