// Package repository defines the repository interfaces for the application.
package repository

import (
	"context"

	"github.com/ezex-io/ezex-users/internal/core/entity"
)

type SecurityImageRepository interface {
	Save(ctx context.Context, image *entity.SecurityImage) error
	GetByID(ctx context.Context, id string) (*entity.SecurityImage, error)
}
