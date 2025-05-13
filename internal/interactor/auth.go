package interactor

import (
	"context"
	"fmt"

	"github.com/ezex-io/ezex-users/internal/entity"
	"github.com/ezex-io/ezex-users/internal/port/repository"
)

type Auth struct {
	userRepo repository.UserRepository
}

func NewAuth(userRepo repository.UserRepository) *Auth {
	return &Auth{
		userRepo: userRepo,
	}
}

func (a *Auth) ProcessFirebaseLogin(
	ctx context.Context,
	req *entity.ProcessFirebaseLoginRequest,
) (*entity.ProcessFirebaseLoginResponse, error) {
	user, err := a.userRepo.GetByEmail(ctx, &entity.GetUserByEmailRequest{
		Email: req.Email,
	})
	if err == nil {
		return &entity.ProcessFirebaseLoginResponse{
			UserID: user.ID,
		}, nil
	}

	newUser, err := a.userRepo.Create(ctx, &entity.CreateUserRequest{
		Email:       req.Email,
		FirebaseUID: req.FirebaseUID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &entity.ProcessFirebaseLoginResponse{
		UserID: newUser.UserID,
	}, nil
}
