package database

import (
	"context"

	"github.com/ezex-io/ezex-users/internal/adapter/database/postgres/gen"
	"github.com/ezex-io/ezex-users/internal/port"
	"github.com/google/uuid"
)

type UserDatabase struct {
	query *gen.Queries
}

func NewUser(query *gen.Queries) *UserDatabase {
	return &UserDatabase{
		query: query,
	}
}

func (u *UserDatabase) CreateUser(ctx context.Context,
	req *port.CreateUserRequest,
) (*port.CreateUserResponse, error) {
	uid := uuid.New()

	if err := u.query.CreateUser(ctx, gen.CreateUserParams{
		ID:           uid,
		Email:        req.Email,
		FirebaseUuid: req.FirebaseUID,
	}); err != nil {
		return nil, err
	}

	return &port.CreateUserResponse{
		UserID: uid.String(),
	}, nil
}

func (u *UserDatabase) GetUserByEmail(ctx context.Context,
	req *port.GetUserByEmailRequest,
) (*port.GetUserByEmailResponse, error) {
	user, err := u.query.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	return &port.GetUserByEmailResponse{
		ID:             user.ID.String(),
		Email:          user.Email,
		FirebaseUID:    user.FirebaseUuid,
		SecurityImage:  user.SecurityImage.String,
		SecurityPhrase: user.SecurityPhrase.String,
	}, nil
}
