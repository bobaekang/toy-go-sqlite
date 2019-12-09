package main

import (
	"log"
)

func main() {
	log.Println("starting...")

	conn := NewSqliteConnection("./toy.db")
	defer conn.Close()

	CreateArrestsTable(conn)

	log.Println("ending...")
}
