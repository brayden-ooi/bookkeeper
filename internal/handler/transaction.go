package handler

import (
	"log"
	"strconv"

	"net/http"

	"github.com/brayden-ooi/bookkeeper/internal/service/account"
	"github.com/brayden-ooi/bookkeeper/internal/service/entry"
	"github.com/brayden-ooi/bookkeeper/internal/service/transaction"
	"github.com/brayden-ooi/bookkeeper/internal/service/user"
	"github.com/brayden-ooi/bookkeeper/internal/view/pages/pages_transactions"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	switch r.Method {
	case http.MethodPost:
		// grab Form
		// year := r.PostFormValue(transaction.Tx_year)
		description := r.PostFormValue(transaction.Tx_description)
		// date := r.PostFormValue(transaction.Tx_date)
		noOfEntriesStr := r.PostFormValue(transaction.Tx_noOfEntries)

		noOfEntries, err := strconv.Atoi(noOfEntriesStr)

		if err != nil {
			log.Fatal("invalid transaction form input")
		}

		counter, err := user.Init(ctx).GetTxCounter()
		if err != nil {
			log.Fatal(err)
		}

		// TODO TEMP FIELDS
		var entry_drafts []entry.Draft

		for i := range noOfEntries {
			tempAccID := r.PostFormValue(entry.AccID(i))
			tempAccType := r.PostFormValue(entry.Type(i))
			tempAccAmt := r.PostFormValue(entry.Amount(i))

			if tempAccID == "" || tempAccType == "" || tempAccAmt == "" {
				log.Fatal("invalid transaction form input")

				return
			}

			entry_drafts = append(entry_drafts, entry.Draft{
				AccountID: tempAccID,
				Type:      tempAccType,
				Amount:    tempAccAmt,
			})
		}

		if _, err := transaction.Init(ctx).UpdateDraft(counter, description, entry_drafts); err != nil {
			log.Fatal(err)
		}

		http.Redirect(w, r, "/transactions", http.StatusSeeOther)
	case http.MethodGet:
		draft, err := transaction.Init(ctx).CreateDraft()

		if err != nil {
			log.Fatal(err)
		}

		// fetch accounts
		accounts, err := account.Init(ctx).List()

		if err != nil {
			log.Fatal(err)
		}

		// create a transaction draft
		pages_transactions.Create(strconv.Itoa(int(draft.Counter)), accounts).Render(ctx, w)
	}
}

func ListTransactions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// fetch transactions data
	txs, err := transaction.Init(ctx).List()

	if err != nil {
		log.Fatal(err)
	}

	// render the page
	pages_transactions.Index(txs).Render(ctx, w)
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	counter := r.URL.Query().Get(transaction.Tx_ID)
	year := r.URL.Query().Get(transaction.Tx_year)
	ctx := r.Context()

	if counter == "" || year == "" {
		log.Fatal("invalid read transaction argument")
	}

	id_int, err := strconv.Atoi(counter)

	if err != nil {
		log.Fatal("invalid read transaction argument")
	}

	tx, entries, err := transaction.Init(ctx).GetByID(id_int)

	if err != nil {
		log.Fatal(err)
	}

	pages_transactions.Edit(tx, entries).Render(ctx, w)
}

// func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
// ctx := r.Context()

// grab Form
// year := r.PostFormValue(transaction.Tx_year)
// description := r.PostFormValue(transaction.Tx_description)
// date := r.PostFormValue(transaction.Tx_date)
// noOfEntriesStr := r.PostFormValue(transaction.Tx_noOfEntries)

// noOfEntries, err := strconv.Atoi(noOfEntriesStr)

// 	if err != nil {
// 		log.Fatal("invalid transaction form input")
// 	}

// 	var entry_drafts []entry.Draft

// 	for i := range noOfEntries {
// 		tempAccID := r.PostFormValue(entry.AccID(i))
// 		tempAccType := r.PostFormValue(entry.Type(i))
// 		tempAccAmt := r.PostFormValue(entry.Amount(i))

// 		if tempAccID == "" || tempAccType == "" || tempAccAmt == "" {
// 			log.Fatal("invalid transaction form input")

// 			return
// 		}

// 		entry_drafts = append(entry_drafts, entry.Draft{
// 			AccountID: tempAccID,
// 			Type:      tempAccType,
// 			Amount:    tempAccAmt,
// 		})
// 	}

// 	http.Redirect(w, r, "/transactions", http.StatusSeeOther)

// }
