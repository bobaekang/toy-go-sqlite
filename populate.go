package main

import (
	"database/sql"
	"log"
)

// PopulateArrestsTable inserts sample rows to Arrests table if missing
func PopulateArrestsTable(db *sql.DB) {
	var count int

	err := db.QueryRow("SELECT Count(*) AS count FROM Arrests").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		addArrestsRow(db, 2019, 1234567)
		addArrestsRow(db, 2018, 1123456)
		addArrestsRow(db, 2017, 1112345)

		log.Println("note: Arrests table populated.")
	}
}

func addArrestsRow(db *sql.DB, year int, value int) {
	stmt, err := db.Prepare("INSERT INTO Arrests(year, value) values(?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(year, value)
	if err != nil {
		log.Fatal(err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("note: id = %d, affected = %d\n", lastID, rowCnt)
}
