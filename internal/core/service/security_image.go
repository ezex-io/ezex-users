package service

import (
	"context"
	"time"

	"github.com/ezex-io/ezex-users/internal/core/entity"
	"github.com/ezex-io/ezex-users/internal/core/model/request"
	"github.com/ezex-io/ezex-users/internal/core/model/response"
	"github.com/ezex-io/ezex-users/internal/core/port/repository"
)

type securityImageService struct {
	repo repository.SecurityImageRepository
}

func NewSecurityImageService(repo repository.SecurityImageRepository) *securityImageService {
	return &securityImageService{
		repo: repo,
	}
}

func (s *securityImageService) SaveSecurityImage(ctx context.Context, req *request.SaveSecurityImageRequest) (*response.SaveSecurityImageResponse, error) {
	image := &entity.SecurityImage{
		ID:        generateID(),
		UserID:    req.UserID,
		ImageData: req.ImageData,
		Metadata:  "foo",
	}

	if err := s.repo.Save(ctx, image); err != nil {
		return nil, err
	}

	return &response.SaveSecurityImageResponse{
		ImageID: image.ID,
	}, nil
}

func (s *securityImageService) GetSecurityImage(ctx context.Context, req *request.GetSecurityImageRequest) (*response.GetSecurityImageResponse, error) {
	image, err := s.repo.GetByID(ctx, req.ImageID)
	if err != nil {
		return nil, err
	}

	return &response.GetSecurityImageResponse{
		ImageData: image.ImageData,
		Metadata:  image.Metadata,
	}, nil
}

func generateID() string {
	return time.Now().Format("20060102150405")
}
