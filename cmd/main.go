package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Kei-K23/go-blog-app/handlers"
	"github.com/Kei-K23/go-blog-app/lib"
	"github.com/Kei-K23/go-blog-app/services"
	"github.com/Kei-K23/go-blog-app/views/errorShow"
	"github.com/Kei-K23/go-blog-app/views/notfound"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var db *sql.DB

func init() {

	var err error

	 err = godotenv.Load(".env")
 if err != nil{
  log.Fatalf("Error loading .env file: %s", err)
 }
 	url := os.Getenv("DB_URL")

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
	
	e.GET("/blogs/:id/edit", func(c echo.Context) error {
		id := c.Param("id")
		blogService := services.BlogService{}
		blog, err := blogService.GetBlog(db, id)

		if err != nil { 
			return handlers.Render(c , notfound.NotfoundShow())
		}
		
		return handlers.Render(c , blogHandler.EditBlog(db, blog))
	})
	
	e.GET("/blogs/:id/delete", func(c echo.Context) error {
		id := c.Param("id")
	
		blogService := services.BlogService{}
		err := blogService.DeleteBlog(db , id)

if  err != nil {
		return handlers.Render(c , errorShow.ErrorShow(err.Error()))
}

		return handlers.Render(c, baseHandler.ShowHome(db))
	})


	e.POST("/blogs/edit", func(c echo.Context) error {
		return blogHandler.UpdateBlog(c , db)
	} )
	

	e.POST("/upload", imageUploadHandler.UploadHandler)
	
	e.GET("/images/*", imageUploadHandler.ServeImages)
	
	e.Logger.Fatal(e.Start(":8080"))

}
