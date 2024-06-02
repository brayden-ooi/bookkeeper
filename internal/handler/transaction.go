package handler

import (
	"log"
	"strconv"

	"net/http"

	"github.com/brayden-ooi/bookkeeper/internal/database"
	"github.com/brayden-ooi/bookkeeper/internal/service/transaction"
	"github.com/brayden-ooi/bookkeeper/internal/view/pages/pages_transactions"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	switch r.Method {
	case http.MethodPost:
		// grab Form
		// TODO

		if _, err := transaction.Init(ctx).Create(); err != nil {
			log.Fatal(err)
		}

		http.Redirect(w, r, "/transactions", http.StatusSeeOther)
	case http.MethodGet:
		pages_transactions.Create().Render(ctx, w)
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

	tx, err := transaction.Init(ctx).GetByID(id_int)

	if err != nil {
		log.Fatal(err)
	}

	pages_transactions.Index([]database.Transaction{tx}).Render(ctx, w)
}
