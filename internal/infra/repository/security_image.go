package repository

import (
	"context"

	"github.com/ezex-io/ezex-users/internal/core/entity"
	"github.com/ezex-io/ezex-users/internal/core/port/repository"
)

type securityImageRepository struct {
}

func NewSecurityImageRepository() repository.SecurityImageRepository {
	return &securityImageRepository{}
}

func (r *securityImageRepository) Save(ctx context.Context, image *entity.SecurityImage) error {
	return nil
}

func (r *securityImageRepository) GetByID(ctx context.Context, id string) (*entity.SecurityImage, error) {
	return &entity.SecurityImage{
		ID:        id,
		Metadata:  "foo",
		ImageData: []byte{},
	}, nil
}
