package postgres

import (
	"context"

	"github.com/ezex-io/ezex-users/internal/adapters/db/postgres/gen"
	"github.com/ezex-io/ezex-users/internal/entity/user"
	"github.com/ezex-io/ezex-users/internal/ports/repository"
)

var _ repository.UserPort = (*userRepo)(nil)

type userRepo struct {
	q gen.Querier
}

func newUserRepo(q gen.Querier) repository.UserPort {
	return &userRepo{
		q: q,
	}
}

func (u *userRepo) GetUserByID(ctx context.Context, id string) (*user.User, error) {
	usr, err := u.q.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return userToDomain(&usr), nil
}
