package repository

import (
	"context"

	"github.com/ezex-io/ezex-users/internal/entity/user"
)

type UserPort interface {
	GetUserByID(ctx context.Context, id string) (*user.User, error)
}
