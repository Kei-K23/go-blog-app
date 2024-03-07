package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Kei-K23/go-blog-app/handlers"
	"github.com/Kei-K23/go-blog-app/lib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var db *sql.DB

func init() {

	var err error

	url := "libsql://go-blog-app-kei-k23.turso.io?authToken=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MDk3MjI4MzEsImlkIjoiZmQ0MmYxMTgtNjQ0Mi00ZWZlLWEwODMtZmU5ODg5OTUyMTRkIn0.1N0st4ziWW-w9EJdyktWX1sLgXQ-mCk2g2f3zYeE5aLmb8lveHusvWJPaVC1t4trwBq4vf-Lp65s-S42YEoSBA"

	db, err = sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}

	dbLib := lib.DBLib{}
	dbLib.CreateBlogTable(db)
	fmt.Println("Successfully connected to database")
}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	baseHandler := handlers.BaseHandler{}
	blogHandler := handlers.BlogHandler{}
	imageUploadHandler := handlers.ImagesUploadHandler{}


	e.GET("/", func(c echo.Context) error {
		return handlers.Render(c, baseHandler.ShowHome(db))
	})

	e.GET("/blogs/:id", func(c echo.Context) error {
		id := c.Param("id")

		return handlers.Render(c, blogHandler.ShowBlog(db, id))
	})
	

	e.GET("/blogs" , func(c echo.Context) error {
		return handlers.Render(c , blogHandler.ShowCreate(db))
	})

	e.POST("/blogs", func(c echo.Context) error {
		return blogHandler.CreateBlog(c, db)
	})

	e.POST("/upload", imageUploadHandler.UploadHandler)
	
	e.GET("/images/*", imageUploadHandler.ServeImages)
	
	e.Logger.Fatal(e.Start(":8080"))

}
