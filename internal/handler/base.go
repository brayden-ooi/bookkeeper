package handler

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	page "github.com/brayden-ooi/bookkeeper/internal/view/pages"
)

func Base(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		page.NotFound().Render(context.Background(), w)

		return
	}
	page.Index().Render(context.Background(), w)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong!"))
}

func NotFound() *templ.ComponentHandler {
	return templ.Handler(page.Index())
}
