// Package service defines the service interfaces for the application.
package service

import (
	"context"

	"github.com/ezex-io/ezex-users/internal/core/model/request"
	"github.com/ezex-io/ezex-users/internal/core/model/response"
)

// SecurityImageService defines the interface for security image operations.
type SecurityImageService interface {
	// SaveSecurityImage saves a security image and returns its ID.
	SaveSecurityImage(
		ctx context.Context,
		req *request.SaveSecurityImageRequest,
	) (*response.SaveSecurityImageResponse, error)

	// GetSecurityImage retrieves a security image by ID.
	GetSecurityImage(
		ctx context.Context,
		req *request.GetSecurityImageRequest,
	) (*response.GetSecurityImageResponse, error)
}
