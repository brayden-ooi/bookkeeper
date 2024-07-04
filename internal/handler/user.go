package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/brayden-ooi/bookkeeper/internal/service/user"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// validate form input
	email := r.FormValue(user.Email)
	pw := r.FormValue(user.Password)
	confirmPw := r.FormValue(user.ConfirmPassword)

	if email == "" || pw != confirmPw {
		w.WriteHeader(http.StatusUnprocessableEntity)

		return
	}

	// create user
	id, err := user.Create(ctx, "admin", email, pw)

	if err != nil {
		log.Fatal("invalid create user argument: ", err)
	}

	// update the context to attach the user id

	// after everything setup redirect user to dashboard
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    strconv.Itoa(id),
		HttpOnly: true,
	})
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	// TODO

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
