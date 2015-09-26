package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/haibin/bookstore/models"
)

type Env struct {
	db *sql.DB
}

func NewEnv(db *sql.DB) Env {
	return Env{db: db}
}

func (env Env) BooksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	bks, err := models.AllBooks(env.db)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
