package server

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	securityimagev1 "github.com/ezex-io/ezex-users/api/gen/go/security_image/v1"
	"github.com/ezex-io/ezex-users/internal/core/port/service"
)

type SecurityImageServer struct {
	securityimagev1.UnimplementedSecurityImageServiceServer
	service service.SecurityImageService
}

func NewSecurityImageServer(service service.SecurityImageService) *SecurityImageServer {
	return &SecurityImageServer{
		service: service,
	}
}

func (*SecurityImageServer) SaveSecurityImage(
	_ context.Context,
	_ *securityimagev1.SaveSecurityImageRequest,
) (*securityimagev1.SaveSecurityImageResponse, error) {
	return &securityimagev1.SaveSecurityImageResponse{
		ImageId: "image-id",
	}, nil
}

func (*SecurityImageServer) GetSecurityImage(
	_ context.Context,
	_ *securityimagev1.GetSecurityImageRequest,
) (*securityimagev1.GetSecurityImageResponse, error) {
	// return moon.png and return with "foo" metadata
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %w", err)
	}

	imagePath := filepath.Join(wd, "moon.png")
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read image file: %w", err)
	}

	return &securityimagev1.GetSecurityImageResponse{
		ImageData: imageData,
		Metadata:  "foo",
	}, nil
}
