package port

import (
	"context"
)

type CreateUserRequest struct {
	FirebaseUID string
	Email       string
}

type CreateUserResponse struct {
	UserID string
}

type GetUserByEmailRequest struct {
	Email string
}

type GetUserByEmailResponse struct {
	ID             string
	Email          string
	FirebaseUID    string
	SecurityImage  string
	SecurityPhrase string
}

type SaveSecurityImageRequest struct {
	Email          string
	SecurityImage  string
	SecurityPhrase string
}

type SaveSecurityImageResponse struct {
	Email string
}

type GetSecurityImageRequest struct {
	Email string
}

type GetSecurityImageResponse struct {
	SecurityImage  string
	SecurityPhrase string
}

type UserDatabasePort interface {
	CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error)
	GetUserByEmail(ctx context.Context, req *GetUserByEmailRequest) (*GetUserByEmailResponse, error)
}

type SecurityImageDatabasePort interface {
	SaveSecurityImage(ctx context.Context, image *SaveSecurityImageRequest) (*SaveSecurityImageResponse, error)
	GetSecurityImage(ctx context.Context, req *GetSecurityImageRequest) (*GetSecurityImageResponse, error)
}
