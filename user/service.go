package user

import (
	"context"
	"example/domain"

	"github.com/samber/do"
)

type service struct {
	repo domain.UserRepository
}

func (s *service) FetchByUsername(ctx context.Context, username string) (*domain.User, error) {
	userEntity, err := s.repo.FetchByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		ID:       userEntity.ID,
		Username: userEntity.Username,
	}, nil
}

// be care that the return type should be interface, rather than the concrete type
func NewService(i *do.Injector) (domain.UserService, error) {
	repo := do.MustInvoke[domain.UserRepository](i)
	return &service{repo: repo}, nil
}
