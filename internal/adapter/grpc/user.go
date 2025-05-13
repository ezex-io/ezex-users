package grpc

import (
	"github.com/ezex-io/ezex-proto/go/users"
	"github.com/ezex-io/ezex-users/internal/interactor"
)

type UserServer struct {
	users.UnimplementedUsersServiceServer

	securityImage *interactor.SecurityImage
	auth          *interactor.Auth
}

func NewUserServer(securityImage *interactor.SecurityImage, auth *interactor.Auth) *UserServer {
	return &UserServer{
		securityImage: securityImage,
		auth:          auth,
	}
}
