package controller

import (
	"github.com/ezex-io/ezex-users/internal/core/port/service"
)

type SecurityImageController struct {
	service service.SecurityImageService
}

func NewSecurityImageController(service service.SecurityImageService) *SecurityImageController {
	return &SecurityImageController{
		service: service,
	}
}
