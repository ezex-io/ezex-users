package grpc

import (
	"context"
	"fmt"

	"github.com/ezex-io/ezex-proto/go/users"
	"github.com/ezex-io/ezex-users/internal/entity"
)

func (s *UserServer) SaveSecurityImage(
	ctx context.Context,
	req *users.SaveSecurityImageRequest,
) (*users.SaveSecurityImageResponse, error) {
	err := s.securityImage.SaveSecurityImage(ctx, &entity.SaveSecurityImageRequest{
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
	resp, err := s.securityImage.GetSecurityImage(ctx, &entity.GetSecurityImageRequest{
		Email: req.Email,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get security image: %w", err)
	}

	return &users.GetSecurityImageResponse{
		SecurityImage:  resp.SecurityImage,
		SecurityPhrase: resp.SecurityPhrase,
	}, nil
}
