// Package response provides response models for the application.
package response

type SaveSecurityImageResponse struct{}

type GetSecurityImageResponse struct {
	UserID         string
	SecurityImage  string
	SecurityPhrase string
}
