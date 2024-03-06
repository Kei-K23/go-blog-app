package handlers

import (
	"database/sql"
	"fmt"

	"github.com/Kei-K23/go-blog-app/services"
	"github.com/Kei-K23/go-blog-app/views/home"
	"github.com/a-h/templ"
)

type BaseHandler struct{}

func (h *BaseHandler) ShowHome(db *sql.DB) templ.Component {
	blogService := services.BlogService{}
	blogs , err := blogService.GetBlogs(db)

	if err != nil {
		fmt.Println("Error getting blog")
	}

	return home.ShowHome(blogs)
}
