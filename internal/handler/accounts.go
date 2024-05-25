package handler

import (
	"log"
	"net/http"

	"github.com/brayden-ooi/bookkeeper/internal/database"
	"github.com/brayden-ooi/bookkeeper/internal/service/account"
	"github.com/brayden-ooi/bookkeeper/internal/view/pages"
)

// func CreateAccount() {

// }

func ListAccounts(w http.ResponseWriter, r *http.Request) {
	_id := r.URL.Query().Get("id")
	ctx := r.Context()

	if _id != "" {
		acc, err := account.GetByID(ctx, _id)

		if err != nil {
			log.Fatal(err)
		}

		pages.Accounts([]database.Account{acc}).Render(ctx, w)
	} else {
		// fetch accounts data
		accounts, err := account.List(ctx)

		if err != nil {
			log.Fatal(err)
		}

		// render the page
		pages.Accounts(accounts).Render(ctx, w)
	}
}
