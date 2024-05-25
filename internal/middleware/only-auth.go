package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/brayden-ooi/bookkeeper/internal/utils"
)

func OnlyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check if the jwt is valid
		fmt.Println("user is authed!")

		ctx := context.WithValue(r.Context(), utils.CtxKeyID, 1)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
