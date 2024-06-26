package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/brayden-ooi/bookkeeper/internal/utils"
)

func OnlyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if r.URL.Path != "/sign-in" && r.URL.Path != "/sign-up" {
			cookie, err := r.Cookie("auth_token")

			if err != nil {
				switch {
				case errors.Is(err, http.ErrNoCookie):
					http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
				default:
					log.Println(err)
					http.Error(w, "server error", http.StatusInternalServerError)
				}
				return
			}

			tokenStr := cookie.Value
			token, err := strconv.Atoi(tokenStr)

			if err != nil {
				log.Fatal("invalid auth_token")
				http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
				return
			}

			ctx = context.WithValue(r.Context(), utils.CtxKeyID, token)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
