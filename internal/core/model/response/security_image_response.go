package response

type SaveSecurityImageResponse struct {
	ImageID string
}

type GetSecurityImageResponse struct {
	ImageData []byte
	Metadata  string
}
