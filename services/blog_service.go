package services

import (
	"database/sql"
	"time"
)

type Blog struct {
	ID int 
	Title string
	Description string
	Body string
	IsPublic bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BlogService struct { }

func (s *BlogService ) CreateBlog(db *sql.DB, blog Blog) error {
	_ , err := db.Exec("INSERT INTO blogs (title, description, body, is_public, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", blog.Title, blog.Description, blog.Body, blog.IsPublic, blog.CreatedAt, blog.UpdatedAt)
	return err
}