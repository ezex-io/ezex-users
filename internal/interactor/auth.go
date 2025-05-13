package interactor

import (
	"context"
	"fmt"

	"github.com/ezex-io/ezex-users/internal/port"
)

type Auth struct {
	userDB port.UserDatabasePort
}

func NewAuth(userDB port.UserDatabasePort) *Auth {
	return &Auth{
		userDB: userDB,
	}
}

func (a *Auth) ProcessLogin(
	ctx context.Context,
	req *port.ProcessLoginRequest,
) (*port.ProcessLoginResponse, error) {
	user, err := a.userDB.GetUserByEmail(ctx, &port.GetUserByEmailRequest{
		Email: req.Email,
	})
	if err == nil {
		return &port.ProcessLoginResponse{
			UserID: user.ID,
		}, nil
	}

	newUser, err := a.userDB.CreateUser(ctx, &port.CreateUserRequest{
		Email:       req.Email,
		FirebaseUID: req.FirebaseUID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &port.ProcessLoginResponse{
		UserID: newUser.UserID,
	}, nil
}
