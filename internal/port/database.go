package port

import (
	"context"

	"github.com/ezex-io/ezex-proto/go/users"
)

type UsersDBPort interface {
	CreateUser(ctx context.Context, req *users.CreateUserRequest) (*users.CreateUserResponse, error)
	GetUserByEmail(ctx context.Context, req *users.GetUserByEmailRequest) (*users.GetUserByEmailResponse, error)
}

type SecurityImageDBPort interface {
	SaveSecurityImage(ctx context.Context, image *users.SaveSecurityImageRequest) (*users.SaveSecurityImageResponse, error)
	GetSecurityImage(ctx context.Context, req *users.GetSecurityImageRequest) (*users.GetSecurityImageResponse, error)
}
