package database

import (
	"context"

	"github.com/ezex-io/ezex-proto/go/users"
	"github.com/ezex-io/ezex-users/internal/adapter/database/postgres/gen"
	"github.com/jackc/pgx/v5/pgtype"
)

type SecurityImage struct {
	query *gen.Queries
}

func NewSecurityImage(query *gen.Queries) *SecurityImage {
	return &SecurityImage{
		query: query,
	}
}

func (s *SecurityImage) SaveSecurityImage(ctx context.Context, req *users.SaveSecurityImageRequest) (
	*users.SaveSecurityImageResponse, error,
) {
	err := s.query.SaveSecurityImage(ctx, gen.SaveSecurityImageParams{
		Email: req.Email,
		SecurityImage: pgtype.Text{
			String: req.SecurityImage,
			Valid:  true,
		},
		SecurityPhrase: pgtype.Text{
			String: req.SecurityPhrase,
			Valid:  true,
		},
	})
	if err != nil {
		return nil, err
	}

	return &users.SaveSecurityImageResponse{
		Email: req.Email,
	}, nil
}

func (s *SecurityImage) GetSecurityImage(ctx context.Context,
	req *users.GetSecurityImageRequest,
) (*users.GetSecurityImageResponse, error) {
	sec, err := s.query.GetSecurityImageByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	return &users.GetSecurityImageResponse{
		SecurityImage:  sec.SecurityImage.String,
		SecurityPhrase: sec.SecurityPhrase.String,
	}, nil
}
