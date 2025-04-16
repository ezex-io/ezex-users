// Package repository defines the interfaces for repository implementations.
package repository

// UserRepository defines the interface for all user-related data access.
type UserRepository interface {
	// SecurityImage returns the security image repository implementation.
	SecurityImage() SecurityImageRepository
}
