package repository

import (
	"context"

	"github.com/ezex-io/ezex-users/internal/adapter/db/postgres/gen"
	"github.com/ezex-io/ezex-users/internal/entity"
	"github.com/ezex-io/ezex-users/internal/port/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type securityImage struct {
	query *gen.Queries
}

func NewSecurityImage(query *gen.Queries) repository.SecurityImageRepository {
	return &securityImage{
		query: query,
	}
}

func (s *securityImage) Save(ctx context.Context, req *entity.SaveSecurityImageRequest) error {
	return s.query.SaveSecurityImage(ctx, gen.SaveSecurityImageParams{
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
}

func (s *securityImage) Get(ctx context.Context,
	req *entity.GetSecurityImageRequest,
) (*entity.GetSecurityImageResponse, error) {
	sec, err := s.query.GetSecurityImageByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	return &entity.GetSecurityImageResponse{
		SecurityImage:  sec.SecurityImage.String,
		SecurityPhrase: sec.SecurityPhrase.String,
	}, nil
}
