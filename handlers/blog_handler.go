package handlers

import (
	"github.com/Kei-K23/go-blog-app/views/create"
	"github.com/a-h/templ"
)

type BlogHandler struct {}

func (h *BlogHandler) ShowCreate() templ.Component {
		return create.ShowCreatePage()
}