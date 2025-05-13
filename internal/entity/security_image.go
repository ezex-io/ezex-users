package entity

type SaveSecurityImageRequest struct {
	Email          string
	SecurityImage  string
	SecurityPhrase string
}

type GetSecurityImageRequest struct {
	Email string
}

type GetSecurityImageResponse struct {
	SecurityImage  string
	SecurityPhrase string
}
