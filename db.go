package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("db ..")

	db, err := sql.Open("postgres", "postgres://postgres:@172.19.0.2/postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
