// Package repository provides data access implementations for the application.
package repository

import (
	"context"

	"github.com/ezex-io/ezex-users/internal/core/entity"
	"github.com/ezex-io/ezex-users/internal/core/port/repository"
)

type securityImageRepository struct{}

func NewSecurityImageRepository() repository.SecurityImageRepository {
	return &securityImageRepository{}
}

func (securityImageRepository) Save(_ context.Context, _ *entity.SecurityImage) error {
	return nil
}

func (securityImageRepository) GetByID(_ context.Context, id string) (*entity.SecurityImage, error) {
	return &entity.SecurityImage{
		UserID:         id,
		SecurityImage:  "moon.png",
		SecurityPhrase: "foo",
	}, nil
}
