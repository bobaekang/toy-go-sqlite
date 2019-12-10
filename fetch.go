package main

import (
	"database/sql"
	"log"
)

// Arrest models a row in Arrests table
type Arrest struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

// FetchArrests fetches all rows in Arrests table
func FetchArrests(db *sql.DB) (arrests []*Arrest) {
	rows, err := db.Query("SELECT * FROM Arrests")
	defer rows.Close()

	for rows.Next() {
		arrest := new(Arrest)

		if err = rows.Scan(&arrest.Year, &arrest.Value); err != nil {
			log.Fatal(err)
		}

		arrests = append(arrests, arrest)
	}

	return arrests
}
