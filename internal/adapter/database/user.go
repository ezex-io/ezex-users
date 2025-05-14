package database

import (
	"context"

	"github.com/ezex-io/ezex-proto/go/users"
	"github.com/ezex-io/ezex-users/internal/adapter/database/postgres/gen"
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
	req *users.CreateUserRequest,
) (*users.CreateUserResponse, error) {
	uid := uuid.New()

	if err := u.query.CreateUser(ctx, gen.CreateUserParams{
		ID:           uid,
		Email:        req.Email,
		FirebaseUuid: req.FirebaseUid,
	}); err != nil {
		return nil, err
	}

	return &users.CreateUserResponse{
		UserId: uid.String(),
	}, nil
}

func (u *UserDatabase) GetUserByEmail(ctx context.Context,
	req *users.GetUserByEmailRequest,
) (*users.GetUserByEmailResponse, error) {
	user, err := u.query.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	return &users.GetUserByEmailResponse{
		Email: user.Email,
	}, nil
}
