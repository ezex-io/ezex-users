// Package request provides request models for the application.
package request

type SaveSecurityImageRequest struct {
	UserID    string
	ImageData []byte
}

type GetSecurityImageRequest struct {
	ImageID string
}
