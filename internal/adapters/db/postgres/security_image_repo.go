package postgres

import (
	"context"
	"errors"
	"sync"

	"github.com/ezex-io/ezex-users/internal/entity"
	"github.com/ezex-io/ezex-users/internal/ports/repository"
)

type securityImageRepo struct {
	lk sync.RWMutex

	mem map[string]*entity.SecurityImage
}

func newSecurityImageRepo() repository.SecurityImagePort {
	return &securityImageRepo{
		mem: make(map[string]*entity.SecurityImage),
	}
}

func (s *securityImageRepo) Save(_ context.Context, image *entity.SecurityImage) error {
	s.lk.RLock()
	defer s.lk.RUnlock()

	s.mem[image.UserID] = image

	return nil
}

func (s *securityImageRepo) GetByID(_ context.Context, id string) (*entity.SecurityImage, error) {
	img, ok := s.mem[id]
	if !ok {
		return nil, errors.New("security image not found")
	}

	return img, nil
}
