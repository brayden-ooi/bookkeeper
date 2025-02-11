// transaction has an ID key but it is only for internal use
// external party (ie the users) will use the counter value or any permutations using the counter as the unique identifier for the transactions
package transaction

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/brayden-ooi/bookkeeper/internal/database"
	"github.com/brayden-ooi/bookkeeper/internal/service"
	"github.com/brayden-ooi/bookkeeper/internal/service/entry"
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
	Tx_date        CreateTxFormKey = "date"
	Tx_noOfEntries CreateTxFormKey = "noOfEntries"
)

type TxExternalIDKey = string

const (
	Tx_year TxExternalIDKey = "year"
	Tx_ID   TxExternalIDKey = "id"
)

type TxIdentifier struct {
	Year int
	ID   int
}

func (srv *tx_service) CreateDraft() (database.Transaction, error) {
	return service.DB.CreateDraft(srv.ctx, database.CreateDraftParams{
		ID:     srv.user_id,
		UserID: srv.user_id,
	})
}

func (srv *tx_service) UpdateDraft(counter int64, description string, entries []entry.Draft, year string, date string) (database.Transaction, error) {
	// TODO refactor to validate entry input first
	yearInt, err := strconv.Atoi(year)

	if err != nil {
		return database.Transaction{}, err
	}

	dateTime, err := time.Parse(time.DateOnly, date)

	if err != nil {
		return database.Transaction{}, err
	}

	tx, err := service.DB.UpdateDraft(srv.ctx, database.UpdateDraftParams{
		Year:        int64(yearInt),
		Description: description,
		Counter:     counter,
		UserID:      srv.user_id,
		Date:        dateTime.Unix(),
	})

	if err != nil {
		return database.Transaction{}, err
	}

	// create entries
	err = entry.Init(srv.ctx).BulkCreate(tx.ID, entries)

	if err != nil {
		return database.Transaction{}, err
	}

	return tx, nil
}

// read operations
func (srv *tx_service) GetByID(counter, year string) (database.Transaction, []database.GetEntriesByTxRow, error) {
	counterInt, err := strconv.Atoi(counter)

	if err != nil {
		log.Fatal("invalid read transaction argument")
	}

	yearInt, err := strconv.Atoi(year)

	if err != nil {
		log.Fatal("invalid read transaction argument")
	}

	tx, err := service.DB.GetTransactionByUserAndID(srv.ctx, database.GetTransactionByUserAndIDParams{
		UserID:  srv.user_id,
		Counter: int64(counterInt),
		Year:    int64(yearInt),
	})

	if err != nil {
		return tx, []database.GetEntriesByTxRow{}, err
	}

	entries, err := entry.Init(srv.ctx).GetByTx(tx.ID)

	return tx, entries, err
}

func (srv *tx_service) List() ([]database.Transaction, error) {
	return service.DB.ListTransactionsByUser(srv.ctx, srv.user_id)
}

func (srv *tx_service) Update(counter, description, yearPrev, yearNext, date string) (database.Transaction, error) {
	counterInt, err := strconv.Atoi(counter)

	if err != nil {
		return database.Transaction{}, err
	}

	yearPrevInt, err := strconv.Atoi(yearPrev)

	if err != nil {
		return database.Transaction{}, err
	}

	yearNextInt, err := strconv.Atoi(yearNext)

	if err != nil {
		return database.Transaction{}, err
	}

	dateTime, err := time.Parse(time.DateOnly, date)

	if err != nil {
		return database.Transaction{}, err
	}

	return service.DB.UpdateTransaction(srv.ctx, database.UpdateTransactionParams{
		Year:        int64(yearNextInt),
		Description: description,
		Status:      "active",
		Date:        dateTime.Unix(),
		Year_2:      int64(yearPrevInt),
		Counter:     int64(counterInt),
		UserID:      srv.user_id,
	})
}
