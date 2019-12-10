package main

import (
	"log"
)

func main() {
	log.Println("starting...")

	conn := NewSqliteConnection("./toy.db")
	defer conn.Close()

	CreateArrestsTable(conn)

	PopulateArrestsTable(conn)

	arrests := FetchArrests(conn)
	log.Println(arrests)

	log.Println("ending...")
}
