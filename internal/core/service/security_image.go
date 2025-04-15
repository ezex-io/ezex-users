// Package service provides the security image service.
package service

import (
	"context"
	"fmt"
	"time"

	"github.com/ezex-io/ezex-users/internal/core/entity"
	"github.com/ezex-io/ezex-users/internal/core/model/request"
	"github.com/ezex-io/ezex-users/internal/core/model/response"
	"github.com/ezex-io/ezex-users/internal/core/port/repository"
)

type SecurityImageService struct {
	repo repository.SecurityImageRepository
}

func NewSecurityImageService(repo repository.SecurityImageRepository) *SecurityImageService {
	return &SecurityImageService{
		repo: repo,
	}
}

func (s *SecurityImageService) SaveSecurityImage(
	ctx context.Context,
	req *request.SaveSecurityImageRequest,
) (*response.SaveSecurityImageResponse, error) {
	image := &entity.SecurityImage{
		ID:        generateID(),
		UserID:    req.UserID,
		ImageData: req.ImageData,
		Metadata:  "foo",
	}

	if err := s.repo.Save(ctx, image); err != nil {
		return nil, fmt.Errorf("failed to save security image: %w", err)
	}

	return &response.SaveSecurityImageResponse{
		ImageID: image.ID,
	}, nil
}

func (s *SecurityImageService) GetSecurityImage(
	ctx context.Context,
	req *request.GetSecurityImageRequest,
) (*response.GetSecurityImageResponse, error) {
	image, err := s.repo.GetByID(ctx, req.ImageID)
	if err != nil {
		return nil, fmt.Errorf("failed to get security image: %w", err)
	}

	return &response.GetSecurityImageResponse{
		ImageData: image.ImageData,
		Metadata:  image.Metadata,
	}, nil
}

func generateID() string {
	return time.Now().Format("20060102150405")
}
