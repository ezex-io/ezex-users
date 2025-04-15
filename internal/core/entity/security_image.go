// Package entity defines the core domain entities of the application.
package entity

type SecurityImage struct {
	ID        string
	UserID    string
	ImageData []byte
	Metadata  string
}
