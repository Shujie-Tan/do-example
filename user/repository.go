package user

import (
	"context"
	"database/sql"
	"example/domain"

	"github.com/samber/do"
)

type repository struct {
	db *sql.DB
}

func (r *repository) FetchByUsername(ctx context.Context, username string) (*domain.UserEntity, error) {
	if username == "joshua" {
		// we mock the data here
		return &domain.UserEntity{
			ID:       "1",
			Username: "joshua",
		}, nil
	}

	// of course we can use db here
	return nil, sql.ErrNoRows
}

func NewRepository(i *do.Injector) (domain.UserRepository, error) {
	db := do.MustInvokeNamed[*sql.DB](i, "user")
	return &repository{db: db}, nil
}
