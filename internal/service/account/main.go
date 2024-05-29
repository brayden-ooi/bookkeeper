package account

import (
	"context"
	"errors"
	"strconv"

	"github.com/brayden-ooi/bookkeeper/internal/database"
	"github.com/brayden-ooi/bookkeeper/internal/service"
	"github.com/brayden-ooi/bookkeeper/internal/utils"
)

type acc_service struct {
	user_id int64
	ctx     context.Context
}

func Init(ctx context.Context) *acc_service {
	return &acc_service{
		user_id: utils.MyContext{ctx}.GetUserID(),
		ctx:     ctx,
	}
}

// write operations
type CreateAccountFormKey = string

const (
	Acc_id   CreateAccountFormKey = "id"
	Acc_name CreateAccountFormKey = "name"
	Acc_type CreateAccountFormKey = "type"
)

func (srv *acc_service) Create(id, name, t string) (database.Account, error) {
	if id == "" || name == "" || t == "" {
		return database.Account{}, errors.New("invalid create account arguments")
	}

	return service.DB.CreateAccount(srv.ctx, database.CreateAccountParams{
		ID:      id,
		Name:    name,
		Type:    t,
		OwnerID: srv.user_id,
	})
}

func (srv *acc_service) DeleteByID(id string) error {
	if id == "" {
		return errors.New("invalid delete account argument")
	}
	return service.DB.DeleteAccountByUserAndID(srv.ctx, database.DeleteAccountByUserAndIDParams{})
}

// read operations
func (srv *acc_service) GetByID(acc_id string) (database.Account, error) {
	// if non-numeric, throw early
	if _, err := strconv.Atoi(acc_id); err != nil {
		return database.Account{}, err
	}

	// fetch individual account data
	account, err := service.DB.GetAccountByUserAndID(srv.ctx, database.GetAccountByUserAndIDParams{
		OwnerID: srv.user_id,
		ID:      acc_id,
	})

	if err != nil {
		return database.Account{}, err
	}

	return account, nil
}

func (srv *acc_service) List() ([]database.Account, error) {
	return service.DB.ListAccountsByUser(srv.ctx, srv.user_id)
}
