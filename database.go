package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteConnection(database string) *sql.DB {
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("note: connection to SQLite database established.")

	return db
}
