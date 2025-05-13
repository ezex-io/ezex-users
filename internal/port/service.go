package port

import (
	"context"
)

type ProcessLoginRequest struct {
	Email       string
	FirebaseUID string
}

type ProcessLoginResponse struct {
	UserID string
}

type ServicePort interface {
	ProcessLogin(ctx context.Context, req *ProcessLoginRequest) (*ProcessLoginResponse, error)
}
