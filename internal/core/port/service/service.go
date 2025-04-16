// Package service defines the interfaces for service implementations.
package service

import (
	"github.com/ezex-io/ezex-users/internal/core/port/repository"
)

// Service defines the interface for all services in the application.
type Service interface {
	User() UserService
}

func NewService(repo repository.Repository) Service {
	return &serviceImpl{
		repo: repo,
	}
}

type serviceImpl struct {
	repo repository.Repository
}

func (s *serviceImpl) User() UserService {
	return NewUserService(s.repo)
}
