package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/Kei-K23/go-blog-app/services"
	detailBlog "github.com/Kei-K23/go-blog-app/views/blog"
	"github.com/Kei-K23/go-blog-app/views/create"
	"github.com/a-h/templ"
)

type BlogHandler struct{}

func (h *BlogHandler) ShowCreate(db *sql.DB) templ.Component {
	return create.ShowCreatePage()
}

func (h *BlogHandler) ShowBlog(db *sql.DB, id int) templ.Component {
  blogService := services.BlogService{}
  blog , err := blogService.GetBlog(db, id)

  if err != nil { 
	fmt.Println("Not Found: ", err)
  }

  return detailBlog.ShowBlog(blog)
}

func (h *BlogHandler) CreateBlog(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// Access form values by their names
	title := r.Form.Get("title")
	description := r.Form.Get("description")
	body := r.Form.Get("body")
	isPublic := r.Form.Get("isPublic") == "on"

	blog := services.Blog{
		Title:       title,
		Description: description,
		Body:        body,
		IsPublic:    isPublic,
		CreatedAt:   time.Now().Local().UTC(),
		UpdatedAt:   time.Now().Local().UTC(),
	}

	blogService := services.BlogService{}

	err = blogService.CreateBlog(db, blog)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	http.Redirect(w, r, "http://localhost:8080/", http.StatusFound)
}
