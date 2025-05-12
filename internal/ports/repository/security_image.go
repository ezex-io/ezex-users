// Package repository defines the repository interfaces for the application.
package repository

import (
	"context"

	"github.com/ezex-io/ezex-users/internal/entity"
)

type SecurityImagePort interface {
	Save(ctx context.Context, image *entity.SecurityImage) error
	GetByID(ctx context.Context, id string) (*entity.SecurityImage, error)
}
