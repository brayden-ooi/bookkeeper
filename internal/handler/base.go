package handler

import (
	"net/http"

	"github.com/a-h/templ"
	page "github.com/brayden-ooi/bookkeeper/internal/view/pages/index"
)

func Base() *templ.ComponentHandler {
	return templ.Handler(page.Index())
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong!"))
}
