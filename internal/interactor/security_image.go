package interactor

import (
	"context"

	"github.com/ezex-io/ezex-users/internal/port"
)

type SecurityImage struct {
	db port.SecurityImageDatabasePort
}

func NewSecurityImage(db port.SecurityImageDatabasePort) *SecurityImage {
	return &SecurityImage{
		db: db,
	}
}

func (u *SecurityImage) SaveSecurityImage(
	ctx context.Context,
	req *port.SaveSecurityImageRequest,
) (*port.SaveSecurityImageResponse, error) {
	return u.db.SaveSecurityImage(ctx, req)
}

func (u *SecurityImage) GetSecurityImage(
	ctx context.Context,
	req *port.GetSecurityImageRequest,
) (*port.GetSecurityImageResponse, error) {
	return u.db.GetSecurityImage(ctx, req)
}
