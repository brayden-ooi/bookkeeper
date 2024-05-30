package handler

import (
	"log"
	"net/http"

	"github.com/brayden-ooi/bookkeeper/internal/service/user"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// create user
	user, err := user.Init(ctx).Create("admin")

	if err != nil {
		log.Fatal("invalid create user argument")
	}

	// update the context to attach the user id

	// on create user, create default internal account tags

	// after everything setup redirect user to dashboard
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// TODO

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
