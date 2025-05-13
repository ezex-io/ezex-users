package repository

import (
	"context"

	"github.com/ezex-io/ezex-users/internal/entity"
)

type UserRepository interface {
	Create(ctx context.Context, req *entity.CreateUserRequest) (*entity.CreateUserResponse, error)
	GetByEmail(ctx context.Context, req *entity.GetUserByEmailRequest) (*entity.GetUserByEmailResponse, error)
}
