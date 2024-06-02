// transaction has an ID key but it is only for internal use
// external party (ie the users) will use the counter value or any permutations using the counter as the unique identifier for the transactions
package transaction

import (
	"context"
	"time"

	"github.com/brayden-ooi/bookkeeper/internal/database"
	"github.com/brayden-ooi/bookkeeper/internal/service"
	"github.com/brayden-ooi/bookkeeper/internal/utils"
)

type tx_service struct {
	user_id int64
	ctx     context.Context
}

func Init(ctx context.Context) *tx_service {
	return &tx_service{
		user_id: utils.MyContext{ctx}.GetUserID(),
		ctx:     ctx,
	}
}

// write operations
type CreateTxFormKey = string

const (
	Tx_description CreateTxFormKey = "description"
)

type TxExternalIDKey = string

const (
	Tx_year TxExternalIDKey = "year"
	Tx_ID   TxExternalIDKey = "id"
)

func (srv *tx_service) Create() (database.Transaction, error) {
	tx, err := service.DB.CreateTransaction(srv.ctx, database.CreateTransactionParams{
		Year:        int64(2024),
		ID:          srv.user_id,
		Description: "test",
		CreatedAt:   time.Now().Unix(),
		UserID:      srv.user_id,
	})

	if err != nil {
		return database.Transaction{}, err
	}

	// TODO
	return tx, nil
}

// read operations
func (srv *tx_service) GetByID(counter int) (database.Transaction, error) {
	return service.DB.GetTransactionByUserAndID(srv.ctx, database.GetTransactionByUserAndIDParams{
		UserID:  srv.user_id,
		Counter: int64(counter),
		Year:    int64(2024),
	})
}

func (srv *tx_service) List() ([]database.Transaction, error) {
	return service.DB.ListTransactionsByUser(srv.ctx, srv.user_id)
}
