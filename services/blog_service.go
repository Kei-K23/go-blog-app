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

func (s *BlogService ) GetBlogs(db *sql.DB) ([]Blog, error) {
	rows , err := db.Query("SELECT * FROM blogs ORDER BY created_at DESC")

	if err!= nil {
        return nil, err
    }
	defer rows.Close()
	
	var blogs []Blog

	for rows.Next() { 
		var blog Blog
        err = rows.Scan(&blog.ID, &blog.Title, &blog.Description, &blog.Body, &blog.IsPublic, &blog.CreatedAt, &blog.UpdatedAt)
        if err!= nil {
            return nil, err
        }
        blogs = append(blogs, blog)
	}

	if err := rows.Err() ; err != nil {
		return nil, err
	}

	return blogs , err

}

func (s *BlogService) GetBlog(db *sql.DB, id int) (Blog, error) {
    rows, err := db.Query("SELECT * FROM blogs WHERE id = ?", id)
    if err != nil {
        return Blog{}, err // Return an empty instance of Blog and the error
    }
    defer rows.Close()

    var blog Blog

    for rows.Next() {
        err := rows.Scan(&blog.ID, &blog.Title, &blog.Description, &blog.Body, &blog.IsPublic, &blog.CreatedAt, &blog.UpdatedAt)
        if err != nil {
            return Blog{}, err // Return an empty instance of Blog and the error
        }
    }

    if err := rows.Err(); err != nil {
        return Blog{}, err // Return an empty instance of Blog and the error
    }

    return blog, nil // Return the retrieved blog and no error
}
