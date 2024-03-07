package handlers

import (
	"database/sql"

	"github.com/Kei-K23/go-blog-app/services"
	"github.com/Kei-K23/go-blog-app/views/errorShow"
	"github.com/Kei-K23/go-blog-app/views/home"
	"github.com/a-h/templ"
)

type BaseHandler struct{}

func (h *BaseHandler) ShowHome(db *sql.DB) templ.Component {
	blogService := services.BlogService{}
	blogs, err := blogService.GetBlogs(db)

	if err != nil {
		return errorShow.ErrorShow(err.Error())
	}

	return home.ShowHome(blogs)
}
