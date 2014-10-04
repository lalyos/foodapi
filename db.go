package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("db ..")

	dbUrl := os.Getenv("DBURL")
	if dbUrl == "" {
		panic("Please set the DBURL env variable: postgres://user:pwd@host/dbname?sslmode=disable")
	}

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
