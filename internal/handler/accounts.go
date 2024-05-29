package handler

import (
	"log"
	"net/http"

	"github.com/brayden-ooi/bookkeeper/internal/database"
	"github.com/brayden-ooi/bookkeeper/internal/service/account"
	"github.com/brayden-ooi/bookkeeper/internal/view/pages/pages_accounts"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	switch r.Method {
	case http.MethodPost:
		// grab Form
		acc_id := r.FormValue(account.Acc_id)
		acc_name := r.FormValue(account.Acc_name)
		acc_type := r.FormValue(account.Acc_type)

		if _, err := account.Init(ctx).Create(acc_id, acc_name, acc_type); err != nil {
			log.Fatal(err)
		}

		http.Redirect(w, r, "/accounts", http.StatusSeeOther)
	case http.MethodGet:
		pages_accounts.Create().Render(ctx, w)
	}
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/accounts", http.StatusSeeOther)
}

func ListAccounts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// fetch accounts data
	accounts, err := account.Init(ctx).List()

	if err != nil {
		log.Fatal(err)
	}

	// render the page
	pages_accounts.Index(accounts).Render(ctx, w)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	ctx := r.Context()

	if id == "" {
		log.Fatal("invalid read account argument")
	}

	acc, err := account.Init(ctx).GetByID(id)

	if err != nil {
		log.Fatal(err)
	}

	pages_accounts.Index([]database.Account{acc}).Render(ctx, w)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	ctx := r.Context()

	if id == "" {
		log.Fatal("invalid delete account argument")
	}

	if err := account.Init(ctx).DeleteByID(id); err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/accounts", http.StatusSeeOther)
}
