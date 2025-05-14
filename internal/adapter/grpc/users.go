package grpc

import (
	"context"

	"github.com/ezex-io/ezex-proto/go/users"
	"github.com/ezex-io/ezex-users/internal/interactor"
)

type UsersServer struct {
	users.UnimplementedUsersServiceServer

	securityImage *interactor.SecurityImage
	users         *interactor.Users
}

func NewUsersService(securityImage *interactor.SecurityImage, users *interactor.Users) *UsersServer {
	return &UsersServer{
		securityImage: securityImage,
		users:         users,
	}
}

func (s *UsersServer) SaveSecurityImage(
	ctx context.Context,
	req *users.SaveSecurityImageRequest,
) (*users.SaveSecurityImageResponse, error) {
	return s.securityImage.SaveSecurityImage(ctx, req)
}

func (s *UsersServer) GetSecurityImage(
	ctx context.Context,
	req *users.GetSecurityImageRequest,
) (*users.GetSecurityImageResponse, error) {
	return s.securityImage.GetSecurityImage(ctx, req)
}

func (s *UsersServer) CreateUser(ctx context.Context,
	req *users.CreateUserRequest,
) (*users.CreateUserResponse, error) {
	return s.users.CreateUser(ctx, req)
}

func (s *UsersServer) GetUserByEmail(ctx context.Context,
	req *users.GetUserByEmailRequest,
) (*users.GetUserByEmailResponse, error) {
	return s.users.GetUserByEmail(ctx, req)
}
