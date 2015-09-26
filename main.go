package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/haibin/bookstore/models"
	_ "github.com/lib/pq"
)

// TODO: remove the global
var db *sql.DB

func init() {
	var err error
	// a pool of connections; safe for concurrent access
	db, err = sql.Open("postgres", "postgres://haibin@localhost/bookstore?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	models.InitDB("postgres://haibin@localhost/bookstore?sslmode=disable")

	http.HandleFunc("/books", booksIndex)
	http.ListenAndServe(":3000", nil)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	bks, err := models.AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
