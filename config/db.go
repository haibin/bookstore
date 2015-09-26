package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	return db, nil
}
