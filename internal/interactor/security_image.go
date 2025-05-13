package interactor

import (
	"context"

	"github.com/ezex-io/ezex-users/internal/entity"
	"github.com/ezex-io/ezex-users/internal/port/repository"
)

type SecurityImage struct {
	repo repository.SecurityImageRepository
}

func NewSecurityImage(repo repository.SecurityImageRepository) *SecurityImage {
	return &SecurityImage{
		repo: repo,
	}
}

func (u *SecurityImage) SaveSecurityImage(
	ctx context.Context,
	req *entity.SaveSecurityImageRequest,
) error {
	return u.repo.Save(ctx, req)
}

func (u *SecurityImage) GetSecurityImage(
	ctx context.Context,
	req *entity.GetSecurityImageRequest,
) (*entity.GetSecurityImageResponse, error) {
	return u.repo.Get(ctx, req)
}
