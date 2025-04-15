package entity

type SecurityImage struct {
	ID        string
	UserID    string
	ImageData []byte
	Metadata  string
}
