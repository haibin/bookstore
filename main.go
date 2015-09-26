package main

import (
	"log"
	"net/http"

	"github.com/haibin/bookstore/config"
	"github.com/haibin/bookstore/handlers"
	_ "github.com/lib/pq"
)

func main() {
	db, err := config.NewDB("postgres://haibin@localhost/bookstore?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}

	env := handlers.NewEnv(db)

	http.HandleFunc("/books", env.BooksIndex)
	http.ListenAndServe(":3000", nil)
}
