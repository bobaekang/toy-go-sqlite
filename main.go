package main

import (
	"encoding/json"
	"log"
)

func main() {
	log.Println("starting...")

	conn := NewSqliteConnection("./toy.db")
	defer conn.Close()

	CreateArrestsTable(conn)

	PopulateArrestsTable(conn)

	arrests := FetchArrests(conn)

	arrestsJSON, err := json.Marshal(arrests)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Query result: %s\n", string(arrestsJSON))

	log.Println("ending...")
}
