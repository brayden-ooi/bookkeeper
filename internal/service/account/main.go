package account

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/brayden-ooi/bookkeeper/internal/database"
	"github.com/brayden-ooi/bookkeeper/internal/service"
	"github.com/brayden-ooi/bookkeeper/internal/utils"
)

func Create(ctx context.Context) {
	account, err := service.DB.CreateAccount(ctx, database.CreateAccountParams{
		Name:    "Assets",
		Type:    "debit",
		OwnerID: int64(1),
	})

	fmt.Println(account)

	if err != nil {
		log.Fatal("Cant connect to database:", err)
	}
}

func GetByID(ctx context.Context, acc_id string) (database.Account, error) {
	id, err := strconv.Atoi(acc_id)

	if err != nil {
		return database.Account{}, err
	}

	// fetch individual account data
	account, err := service.DB.GetAccountByUserAndID(ctx, database.GetAccountByUserAndIDParams{
		OwnerID: utils.MyContext{ctx}.GetUserID(),
		ID:      int64(id),
	})

	if err != nil {
		return database.Account{}, err
	}

	return account, nil
}

func List(ctx context.Context) ([]database.Account, error) {
	return service.DB.ListAccountsByUser(ctx, utils.MyContext{ctx}.GetUserID())
}
