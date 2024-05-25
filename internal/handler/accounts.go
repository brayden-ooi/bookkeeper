package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/brayden-ooi/bookkeeper/internal/service"
	"github.com/brayden-ooi/bookkeeper/internal/view/pages"
)

// func CreateAccount() {

// }

func ListAccounts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	// fetch accounts data
	accounts, err := service.DB.ListAccountsByUser(ctx, 1)
	if err != nil {
		log.Fatal("error retrieving accounts")
	}

	// render the page
	pages.Accounts(accounts).Render(ctx, w)
}
