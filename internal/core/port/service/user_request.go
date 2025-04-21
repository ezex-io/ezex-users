package service

type SaveSecurityImageRequest struct {
	UserID         string
	SecurityImage  string
	SecurityPhrase string
}

type GetSecurityImageRequest struct {
	UserID string
}
