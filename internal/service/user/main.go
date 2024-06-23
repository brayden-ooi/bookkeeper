package user

import (
	"context"

	"github.com/brayden-ooi/bookkeeper/internal/database"
	"github.com/brayden-ooi/bookkeeper/internal/service"
	"github.com/brayden-ooi/bookkeeper/internal/utils"
)

type user_service struct {
	user_id int64
	ctx     context.Context
}

func Init(ctx context.Context) *user_service {
	return &user_service{
		user_id: utils.MyContext{ctx}.GetUserID(),
		ctx:     ctx,
	}
}

func Create(ctx context.Context, name string) (database.User, error) {
	return service.DB.CreateUser(ctx, name)
}

// func GetByID(id int) database.User {
// 	ctx := context.Background()

// 	user, err := service.DB.GetUser(ctx, int64(id))

// 	return user
// }

func (srv *user_service) GetTxCounter() (int64, error) {
	return service.DB.GetTransactionCounter(srv.ctx, srv.user_id)
}
