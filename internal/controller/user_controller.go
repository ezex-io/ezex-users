// Package controller provides the controller for the user service.
package controller

import (
	"context"
	"fmt"

	"github.com/ezex-io/ezex-users/internal/core/port/service"
	userspb "github.com/ezex-io/ezex-users/pkg/grpc"
)

type UsersServer struct {
	userspb.UnimplementedUsersServiceServer
	service service.UserService
}

func NewUserServer(service service.UserService) *UsersServer {
	return &UsersServer{
		service: service,
	}
}

func (s *UsersServer) SaveSecurityImage(
	ctx context.Context,
	req *userspb.SaveSecurityImageRequest,
) (*userspb.SaveSecurityImageResponse, error) {
	_, err := s.service.SaveSecurityImage(ctx, &service.SaveSecurityImageRequest{
		UserID:         req.UserId,
		SecurityImage:  req.SecurityImage,
		SecurityPhrase: req.SecurityPhrase,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to save security image: %w", err)
	}

	return &userspb.SaveSecurityImageResponse{}, nil
}

func (s *UsersServer) GetSecurityImage(
	ctx context.Context,
	req *userspb.GetSecurityImageRequest,
) (*userspb.GetSecurityImageResponse, error) {
	resp, err := s.service.GetSecurityImage(ctx, &service.GetSecurityImageRequest{
		UserID: req.UserId,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get security image: %w", err)
	}

	return &userspb.GetSecurityImageResponse{
		UserId:         resp.UserID,
		SecurityImage:  resp.SecurityImage,
		SecurityPhrase: resp.SecurityPhrase,
	}, nil
}
