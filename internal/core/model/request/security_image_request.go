// Package request provides request models for the application.
package request

type SaveSecurityImageRequest struct {
	UserID         string
	SecurityImage  string
	SecurityPhrase string
}

type GetSecurityImageRequest struct {
	UserID string
}
