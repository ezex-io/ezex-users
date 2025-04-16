package server

import (
	"context"
	"fmt"

	"github.com/ezex-io/ezex-users/api/grpc/proto"
	"github.com/ezex-io/ezex-users/internal/core/model/request"
	"github.com/ezex-io/ezex-users/internal/core/port/service"
)

type UserServer struct {
	proto.UnimplementedUserServiceServer
	service service.UserService
}

func NewUserServer(service service.UserService) *UserServer {
	return &UserServer{
		service: service,
	}
}

func (s *UserServer) SaveSecurityImage(
	ctx context.Context,
	req *proto.SaveSecurityImageRequest,
) (*proto.SaveSecurityImageResponse, error) {
	_, err := s.service.SaveSecurityImage(ctx, &request.SaveSecurityImageRequest{
		UserID:         req.UserId,
		SecurityImage:  req.SecurityImage,
		SecurityPhrase: req.SecurityPhrase,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to save security image: %w", err)
	}

	return &proto.SaveSecurityImageResponse{}, nil
}

func (s *UserServer) GetSecurityImage(
	ctx context.Context,
	req *proto.GetSecurityImageRequest,
) (*proto.GetSecurityImageResponse, error) {
	resp, err := s.service.GetSecurityImage(ctx, &request.GetSecurityImageRequest{
		UserID: req.UserId,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get security image: %w", err)
	}

	return &proto.GetSecurityImageResponse{
		UserId:         resp.UserID,
		SecurityImage:  resp.SecurityImage,
		SecurityPhrase: resp.SecurityPhrase,
	}, nil
}
