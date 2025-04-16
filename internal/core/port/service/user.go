// Package service defines the service interfaces for the application.
package service

import (
	"context"

	"github.com/ezex-io/ezex-users/internal/core/model/request"
	"github.com/ezex-io/ezex-users/internal/core/model/response"
)

// UserService defines the interface for user operations.
type UserService interface {
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
