package interactor

import (
	"context"

	"github.com/ezex-io/ezex-proto/go/users"
	"github.com/ezex-io/ezex-users/internal/port"
)

type SecurityImage struct {
	db port.SecurityImageDBPort
}

func NewSecurityImage(db port.SecurityImageDBPort) *SecurityImage {
	return &SecurityImage{
		db: db,
	}
}

func (s *SecurityImage) SaveSecurityImage(
	ctx context.Context,
	req *users.SaveSecurityImageRequest,
) (*users.SaveSecurityImageResponse, error) {
	return s.db.SaveSecurityImage(ctx, req)
}

func (s *SecurityImage) GetSecurityImage(
	ctx context.Context,
	req *users.GetSecurityImageRequest,
) (*users.GetSecurityImageResponse, error) {
	return s.db.GetSecurityImage(ctx, req)
}
