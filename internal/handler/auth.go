package handler

import (
	"github.com/a-h/templ"
	"github.com/brayden-ooi/bookkeeper/internal/view/pages/pages_sign_in"
	"github.com/brayden-ooi/bookkeeper/internal/view/pages/pages_sign_up"
)

func SignUp() *templ.ComponentHandler {
	return templ.Handler(pages_sign_up.Index())
}

func SignIn() *templ.ComponentHandler {
	return templ.Handler(pages_sign_in.Index())
}
