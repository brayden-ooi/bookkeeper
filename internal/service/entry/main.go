package entry

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/brayden-ooi/bookkeeper/internal/database"
	"github.com/brayden-ooi/bookkeeper/internal/service"
	"github.com/brayden-ooi/bookkeeper/internal/utils"
)

type Draft struct {
	AccountID string
	Type      string
	Amount    string
}

type entry_service struct {
	user_id int64
	ctx     context.Context
}

func Init(ctx context.Context) *entry_service {
	return &entry_service{
		user_id: utils.MyContext{ctx}.GetUserID(),
		ctx:     ctx,
	}
}

// write operations
type CreateEntryFormKey = string

func AccID(index int) CreateEntryFormKey {
	return fmt.Sprintf("entry[%d]_acc_id", index)
}

func Type(index int) CreateEntryFormKey {
	return fmt.Sprintf("entry[%d]_type", index)
}

func Amount(index int) CreateEntryFormKey {
	return fmt.Sprintf("entry[%d]_amount", index)
}

func (srv *entry_service) Create(txID int64, draft Draft) (database.Entry, error) {
	// validation
	amt, err := strconv.Atoi(draft.Amount)
	if err != nil {
		return database.Entry{}, err
	}

	if draft.Type != "debit" && draft.Type != "credit" {
		return database.Entry{}, errors.New("invalid entry type")
	}

	if draft.AccountID == "" {
		return database.Entry{}, errors.New("invalid account ID")
	}

	return service.DB.CreateEntry(srv.ctx, database.CreateEntryParams{
		Currency:      "MYR",
		Amount:        int64(amt),
		Type:          draft.Type,
		AccountID:     draft.AccountID,
		AccountUserID: srv.user_id,
		TransactionID: txID,
	})
}

func (srv *entry_service) BulkCreate(txID int64, drafts []Draft) error {
	c := make(chan error)
	done := 0

	for _, entry := range drafts {
		go func(entry Draft) {
			_, err := srv.Create(txID, entry)

			c <- err
		}(entry)
	}

	for {
		err := <-c

		if err != nil {
			return err
		}

		done++

		if done == len(drafts) {
			return nil
		}
	}
}
