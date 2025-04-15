package request

type SaveSecurityImageRequest struct {
	UserID    string
	ImageData []byte
}

type GetSecurityImageRequest struct {
	ImageID string
}
