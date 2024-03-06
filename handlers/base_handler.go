package handlers

import (
	"github.com/Kei-K23/go-blog-app/views/home"
	"github.com/a-h/templ"
)

type BaseHandler struct{}

func (h *BaseHandler) ShowHome() templ.Component {
	return home.ShowHome()
}
