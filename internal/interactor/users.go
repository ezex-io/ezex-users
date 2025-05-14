package interactor

import (
	"context"

	"github.com/ezex-io/ezex-proto/go/users"
	"github.com/ezex-io/ezex-users/internal/port"
)

type Users struct {
	userDB port.UsersDBPort
}

func NewAuth(userDB port.UsersDBPort) *Users {
	return &Users{
		userDB: userDB,
	}
}

func (u *Users) CreateUser(
	ctx context.Context,
	req *users.CreateUserRequest,
) (*users.CreateUserResponse, error) {
	return u.userDB.CreateUser(ctx, req)
}

func (u *Users) GetUserByEmail(
	ctx context.Context,
	req *users.GetUserByEmailRequest,
) (*users.GetUserByEmailResponse, error) {
	return u.userDB.GetUserByEmail(ctx, req)
}
