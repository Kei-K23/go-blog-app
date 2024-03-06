package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Kei-K23/go-blog-app/handlers"
	"github.com/a-h/templ"
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

	fmt.Println("Successfully connected to database")
}

func main() {
	mux := http.NewServeMux()

	baseHandler := handlers.BaseHandler{}
	blogHandler := handlers.BlogHandler{}

	mux.Handle("GET /", templ.Handler(baseHandler.ShowHome(db)))
	mux.HandleFunc("GET /blogs/{id}", func(w http.ResponseWriter, r *http.Request) {
		idPath := r.PathValue("id") 
		id , err := strconv.Atoi(idPath)
		if err!= nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
		templ.Handler(blogHandler.ShowBlog(db, id))
	})

	mux.Handle("GET /create", templ.Handler(blogHandler.ShowCreate(db)))

	mux.HandleFunc("POST /create", func(w http.ResponseWriter, r *http.Request) {
		blogHandler.CreateBlog(w, r , db)
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
