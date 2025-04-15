package service

import (
	"context"

	"github.com/ezex-io/ezex-users/internal/core/model/request"
	"github.com/ezex-io/ezex-users/internal/core/model/response"
)

type SecurityImageService interface {
	SaveSecurityImage(ctx context.Context, req *request.SaveSecurityImageRequest) (*response.SaveSecurityImageResponse, error)
	GetSecurityImage(ctx context.Context, req *request.GetSecurityImageRequest) (*response.GetSecurityImageResponse, error)
}
