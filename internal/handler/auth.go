package handler

import (
	"github.com/a-h/templ"
	"github.com/brayden-ooi/bookkeeper/internal/view/pages"
)

func SignUp() *templ.ComponentHandler {
	return templ.Handler(pages.SignUp())
}

func SignIn() *templ.ComponentHandler {
	return templ.Handler(pages.SignIn())
}
