package main

import (
	"database/sql"
	"log"
)

func CreateArrestsTable(db *sql.DB) {
	if exist := CheckExistArrestsTable(db); !exist {
		_, err := db.Exec("CREATE TABLE Arrests ( year integer, value integer )")
		if err != nil {
			log.Fatal(err)
		}

		log.Println("note: Arrests table created.")
	}
}

func CheckExistArrestsTable(db *sql.DB) (exist bool) {
	err := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='Arrests'")
	if err != nil {
		return false
	} else {
		return true
	}
}
