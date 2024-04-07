package handler

import (
	"github.com/a-h/templ"
	page "github.com/brayden-ooi/bookkeeper/internal/view/pages"
)

func SignUp() *templ.ComponentHandler {
	return templ.Handler(page.SignUp())
}

func SignIn() *templ.ComponentHandler {
	return templ.Handler(page.SignIn())
}
