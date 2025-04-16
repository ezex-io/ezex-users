package server

import (
	"context"

	"github.com/ezex-io/ezex-users/api/grpc/proto"
	"github.com/ezex-io/ezex-users/internal/core/port/service"
)

type UserServer struct {
	proto.UnimplementedUsersServiceServer
	service service.UserService
}

func NewUserServer(service service.UserService) *UserServer {
	return &UserServer{
		service: service,
	}
}

func (*UserServer) SaveSecurityImage(
	_ context.Context,
	_ *proto.SaveSecurityImageRequest,
) (*proto.SaveSecurityImageResponse, error) {
	return &proto.SaveSecurityImageResponse{
		ImageId: "image-id",
	}, nil
}

func (*UserServer) GetSecurityImage(
	_ context.Context,
	_ *proto.GetSecurityImageRequest,
) (*proto.GetSecurityImageResponse, error) {
	return &proto.GetSecurityImageResponse{
		ImageData: []byte("moon.png"),
		Metadata:  "foo",
	}, nil
}
