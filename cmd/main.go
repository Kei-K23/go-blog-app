package main

import (
	"log"
	"net/http"

	"github.com/Kei-K23/go-blog-app/handlers"
	"github.com/a-h/templ"
)

func main()  {
	mux := http.NewServeMux()

 	baseHandler := handlers.BaseHandler{}	
 	blogHandler := handlers.BlogHandler{}	

	mux.Handle("GET /", templ.Handler(baseHandler.ShowHome()))
	mux.Handle("GET /create", templ.Handler(blogHandler.ShowCreate()))

	log.Fatal(http.ListenAndServe(":8080", mux))
}