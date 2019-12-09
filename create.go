package main

import (
	"database/sql"
	"log"
)

func CreateArrestsTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE Arrests ( year integer, value integer )")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("note: Arrests table created.")
}
