// Package repository defines the interfaces for repository implementations.
package repository

// Repository defines the interface for all repositories in the application.
type Repository interface {
	// SecurityImage returns the security image repository implementation.
	SecurityImage() SecurityImageRepository
}
