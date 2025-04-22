// Package service defines the service interfaces for the application.
package service

import (
	"context"
	"fmt"

	"github.com/ezex-io/ezex-users/internal/core/entity"
	"github.com/ezex-io/ezex-users/internal/core/port/repository"
)

// UserService defines the interface for user operations.
type UserService interface {
	// SaveSecurityImage saves a security image for a user.
	SaveSecurityImage(
		ctx context.Context,
		req *SaveSecurityImageRequest,
	) (*SaveSecurityImageResponse, error)

	// GetSecurityImage retrieves a security image by UserID.
	GetSecurityImage(
		ctx context.Context,
		req *GetSecurityImageRequest,
	) (*GetSecurityImageResponse, error)
}

func NewUserService(repo repository.Repository) UserService {
	return &userService{repo: repo}
}

type userService struct {
	repo repository.Repository
}

func (s *userService) SaveSecurityImage(
	ctx context.Context,
	req *SaveSecurityImageRequest,
) (*SaveSecurityImageResponse, error) {
	err := s.repo.SecurityImage().Save(ctx, &entity.SecurityImage{
		UserID:         req.UserID,
		SecurityImage:  req.SecurityImage,
		SecurityPhrase: req.SecurityPhrase,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to save security image: %w", err)
	}

	return &SaveSecurityImageResponse{}, nil
}

func (s *userService) GetSecurityImage(
	ctx context.Context,
	req *GetSecurityImageRequest,
) (*GetSecurityImageResponse, error) {
	image, err := s.repo.SecurityImage().GetByID(ctx, req.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get security image: %w", err)
	}

	return &GetSecurityImageResponse{
		UserID:         image.UserID,
		SecurityImage:  image.SecurityImage,
		SecurityPhrase: image.SecurityPhrase,
	}, nil
}
