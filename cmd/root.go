/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cmp"
	"net/http"
	"os"

	"github.com/brayden-ooi/bookkeeper/internal/handler"
	"github.com/brayden-ooi/bookkeeper/internal/middleware"
	_ "github.com/brayden-ooi/bookkeeper/internal/service"
	"github.com/joho/godotenv"
)

// Main applications for this project.
// The directory name for each application should match the name of the executable you want to have (e.g., /cmd/myapp).

// If you think the code can be imported and used in other projects, then it should live in the /pkg directory.
// If the code is not reusable or if you don't want others to reuse it, put that code in the /internal directory.

func Execute() {
	godotenv.Load()
	mux := http.NewServeMux()

	mux.Handle("GET /static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("web/static/"))))

	/* authed routes */
	mux.Handle("GET /dashboard", handler.Dashboard())

	// accounts routes
	mux.HandleFunc("GET /accounts", handler.ListAccounts)
	mux.HandleFunc("GET /account", handler.GetAccount)
	mux.HandleFunc("PUT /account", handler.UpdateAccount)
	mux.HandleFunc("DELETE /account", handler.DeleteAccount)
	mux.HandleFunc("GET /account/create", handler.CreateAccount)
	mux.HandleFunc("POST /account/create", handler.CreateAccount)

	// account tags routes
	mux.HandleFunc("GET /account-tags", handler.ListTags)
	mux.HandleFunc("GET /account-tag", handler.GetTag)
	// mux.HandleFunc("PUT /account-tag", handler.UpdateAccount)
	// mux.HandleFunc("DELETE /account-tag", handler.DeleteAccount)
	mux.HandleFunc("GET /account-tag/create", handler.CreateTag)
	mux.HandleFunc("POST /account-tag/create", handler.CreateTag)

	// auth routes
	mux.Handle("GET /sign-up", handler.SignUp())
	mux.HandleFunc("POST /sign-up", handler.CreateUser)
	mux.Handle("GET /sign-in", handler.SignIn())
	mux.HandleFunc("POST /sign-in", handler.AuthenticateUser)

	// base routes
	mux.HandleFunc("GET /", handler.Base)
	mux.HandleFunc("GET /ping", handler.Ping)

	server := http.Server{
		Addr:    ":" + cmp.Or(os.Getenv("PORT"), "8000"),
		Handler: middleware.OnlyAuth(mux),
	}

	server.ListenAndServe()
}
