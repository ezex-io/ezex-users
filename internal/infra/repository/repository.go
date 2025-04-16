// Package repository provides data access implementations for the application.
package repository

import (
	"github.com/ezex-io/ezex-users/internal/core/port/repository"
)

// repositoryImpl implements the Repository interface.
type repositoryImpl struct {
	securityImage repository.SecurityImageRepository
}

func NewRepository() repository.Repository {
	return &repositoryImpl{
		securityImage: NewSecurityImageRepository(),
	}
}

func (r *repositoryImpl) SecurityImage() repository.SecurityImageRepository {
	return r.securityImage
}
