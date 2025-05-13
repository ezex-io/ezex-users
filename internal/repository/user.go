package repository

import (
	"context"

	"github.com/ezex-io/ezex-users/internal/adapter/db/postgres/gen"
	"github.com/ezex-io/ezex-users/internal/entity"
	"github.com/ezex-io/ezex-users/internal/port/repository"
	"github.com/google/uuid"
)

type userRepository struct {
	query *gen.Queries
}

func NewUser(query *gen.Queries) repository.UserRepository {
	return &userRepository{
		query: query,
	}
}

func (u *userRepository) Create(ctx context.Context,
	req *entity.CreateUserRequest,
) (*entity.CreateUserResponse, error) {
	uid := uuid.New()

	if err := u.query.CreateUser(ctx, gen.CreateUserParams{
		ID:           uid,
		Email:        req.Email,
		FirebaseUuid: req.FirebaseUID,
	}); err != nil {
		return nil, err
	}

	return &entity.CreateUserResponse{
		UserID: uid.String(),
	}, nil
}

func (u *userRepository) GetByEmail(ctx context.Context,
	req *entity.GetUserByEmailRequest,
) (*entity.GetUserByEmailResponse, error) {
	user, err := u.query.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	return &entity.GetUserByEmailResponse{
		ID:             user.ID.String(),
		Email:          user.Email,
		FirebaseUID:    user.FirebaseUuid,
		SecurityImage:  user.SecurityImage.String,
		SecurityPhrase: user.SecurityPhrase.String,
	}, nil
}
