package main

import (
	"net/http"

	"github.com/haibin/bookstore/config"
	"github.com/haibin/bookstore/handlers"
	_ "github.com/lib/pq"
)

func main() {
	config.InitDB("postgres://haibin@localhost/bookstore?sslmode=disable")

	http.HandleFunc("/books", handlers.BooksIndex)
	http.ListenAndServe(":3000", nil)
}
