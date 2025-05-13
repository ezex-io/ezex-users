package repository

import (
	"context"

	"github.com/ezex-io/ezex-users/internal/entity"
)

type SecurityImageRepository interface {
	Save(ctx context.Context, image *entity.SaveSecurityImageRequest) error
	Get(ctx context.Context, req *entity.GetSecurityImageRequest) (*entity.GetSecurityImageResponse, error)
}
