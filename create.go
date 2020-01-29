package main

import (
	"database/sql"
	"log"
)

// CreateArrestsTable creates table Arrests in db if missing
func CreateArrestsTable(db *sql.DB) {
	if exist := checkExistArrestsTable(db); !exist {
		_, err := db.Exec("CREATE TABLE Arrests ( year integer, value integer )")
		if err != nil {
			log.Fatal(err)
		}

		log.Println("note: Arrests table created.")
	}
}

func checkExistArrestsTable(db *sql.DB) (exist bool) {
	_, err := db.Query("SELECT * FROM Arrests")
	return err == nil
}
