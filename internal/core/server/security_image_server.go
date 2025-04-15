package server

import (
	"context"
	"os"
	"path/filepath"

	security_imagev1 "github.com/ezex-io/ezex-users/api/gen/go/security_image/v1"
	"github.com/ezex-io/ezex-users/internal/core/port/service"
)

type SecurityImageServer struct {
	security_imagev1.UnimplementedSecurityImageServiceServer
	service service.SecurityImageService
}

func NewSecurityImageServer(service service.SecurityImageService) *SecurityImageServer {
	return &SecurityImageServer{
		service: service,
	}
}

func (s *SecurityImageServer) SaveSecurityImage(ctx context.Context, req *security_imagev1.SaveSecurityImageRequest) (*security_imagev1.SaveSecurityImageResponse, error) {
	return &security_imagev1.SaveSecurityImageResponse{
		ImageId: "image-id",
	}, nil

	// resp, err := s.service.SaveSecurityImage(ctx, &dto.SaveSecurityImageRequest{
	// 	UserID:    req.UserId,
	// 	ImageData: req.ImageData,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// return &api.SaveSecurityImageResponse{
	// 	ImageId: resp.ImageID,
	// }, nil
}

func (s *SecurityImageServer) GetSecurityImage(ctx context.Context, req *security_imagev1.GetSecurityImageRequest) (*security_imagev1.GetSecurityImageResponse, error) {
	// return moon.png and return with "foo" metadata
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	imagePath := filepath.Join(wd, "moon.png")
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}

	return &security_imagev1.GetSecurityImageResponse{
		ImageData: imageData,
		Metadata:  "foo",
	}, nil

	// resp, err := s.service.GetSecurityImage(ctx, &dto.GetSecurityImageRequest{
	// 	ImageID: req.ImageId,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// return &api.GetSecurityImageResponse{
	// 	ImageData: resp.ImageData,
	// 	Metadata:  resp.Metadata,
	// }, nil
}
