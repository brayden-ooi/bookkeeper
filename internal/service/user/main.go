package user

import (
	"context"

	"github.com/brayden-ooi/bookkeeper/internal/database"
	"github.com/brayden-ooi/bookkeeper/internal/service"
)

type user_service struct {
	ctx context.Context
}

func Init(ctx context.Context) *user_service {
	return &user_service{
		ctx,
	}
}

func (srv *user_service) Create(name string) (database.User, error) {
	return service.DB.CreateUser(srv.ctx, name)
}
