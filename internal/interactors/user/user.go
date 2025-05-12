package user

import (
	"context"
	"fmt"

	"github.com/ezex-io/ezex-proto/go/users"
	"github.com/ezex-io/ezex-users/internal/entity"
	"github.com/ezex-io/ezex-users/internal/ports/repository"
)

type User struct {
	repo repository.SecurityImagePort
}

func New(repo repository.SecurityImagePort) *User {
	return &User{repo: repo}
}

func (u *User) SaveSecurityImage(
	ctx context.Context,
	req *users.SaveSecurityImageRequest,
) (*users.SaveSecurityImageResponse, error) {
	err := u.repo.Save(ctx, &entity.SecurityImage{
		UserID:         req.UserId,
		SecurityImage:  req.SecurityImage,
		SecurityPhrase: req.SecurityPhrase,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to save user image: %w", err)
	}

	return &users.SaveSecurityImageResponse{}, nil
}

func (u *User) GetSecurityImage(
	ctx context.Context,
	req *users.GetSecurityImageRequest,
) (*users.GetSecurityImageResponse, error) {
	image, err := u.repo.GetByID(ctx, req.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user image: %w", err)
	}

	return &users.GetSecurityImageResponse{
		UserId:         image.UserID,
		SecurityImage:  image.SecurityImage,
		SecurityPhrase: image.SecurityPhrase,
	}, nil
}
