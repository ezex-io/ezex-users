package service

type SaveSecurityImageResponse struct{}

type GetSecurityImageResponse struct {
	UserID         string
	SecurityImage  string
	SecurityPhrase string
}
