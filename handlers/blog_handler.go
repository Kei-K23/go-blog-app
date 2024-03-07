package handlers

import (
	"database/sql"
	"html/template"
	"time"

	"github.com/Kei-K23/go-blog-app/services"
	detailBlog "github.com/Kei-K23/go-blog-app/views/blog"
	"github.com/Kei-K23/go-blog-app/views/create"
	"github.com/Kei-K23/go-blog-app/views/errorShow"
	"github.com/Kei-K23/go-blog-app/views/notfound"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type BlogHandler struct{}

func (h *BlogHandler) ShowCreate(db *sql.DB) templ.Component {
	return create.ShowCreatePage()
}

func (h *BlogHandler) ShowBlog(db *sql.DB, id string) templ.Component {
	blogService := services.BlogService{}
	blog, err := blogService.GetBlog(db, id)

	if blog.ID == "" {
		return notfound.NotfoundShow()
	}
 
	if err != nil {
		return notfound.NotfoundShow()
	}
	return detailBlog.ShowBlog(blog)
}

func (h *BlogHandler) CreateBlog(e  echo.Context, db *sql.DB) error {

	title := e.FormValue("title")
	description := e.FormValue("description")
	body := e.FormValue("body")
	isPublic := e.FormValue("isPublic") == "on"

	blog := services.Blog{
		Title:       title,
		Description: description,
		Body:        template.HTML(body),
		IsPublic:    isPublic,
		CreatedAt:   time.Now().Local().UTC(),
		UpdatedAt:   time.Now().Local().UTC(),
	}

	blogService := services.BlogService{}

	err := blogService.CreateBlog(db, blog)

	if err != nil {
		return Render(e, errorShow.ErrorShow(err.Error()))
	}
	return Render(e, detailBlog.ShowBlog(blog))
}

