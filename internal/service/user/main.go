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

// write operations
type CreateUserFormKey = string

const (
	Email           CreateUserFormKey = "email"
	Password        CreateUserFormKey = "password"
	ConfirmPassword CreateUserFormKey = "confirm-password"
)

func Create(ctx context.Context, name, email, pw string) (int, error) {
	// create user record
	user, err := service.DB.CreateUser(ctx, name)

	if err != nil {
		return 0, err
	}

	// create user credential record
	err = service.DB.CreateUserCredentials(ctx, database.CreateUserCredentialsParams{
		UserID:   user.ID,
		Email:    email,
		Password: pw,
	})

	return int(user.ID), err
}

// func GetByID(id int) database.User {
// 	ctx := context.Background()

// 	user, err := service.DB.GetUser(ctx, int64(id))

// 	return user
// }

func (srv *user_service) GetTxCounter() (int64, error) {
	return service.DB.GetTransactionCounter(srv.ctx, srv.user_id)
}
