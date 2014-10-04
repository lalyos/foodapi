package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func pingDB() {
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func openDB() {
	dbUrl := os.Getenv("DBURL")
	if dbUrl == "" {
		panic("Please set the DBURL env variable: postgres://user:pwd@host/dbname?sslmode=disable")
	}

	db, err = sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
}

func listTables() {

	log.Println("[tables]")
	schema := "public"
	listTables := `
	SELECT table_name
    FROM information_schema.tables
		WHERE table_schema=$1
		ORDER BY table_name;
	`

	rows, err := db.Query(listTables, schema)
	if err != nil {
		log.Fatal(err)
	}
	var table string
	for rows.Next() {
		err := rows.Scan(&table)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(table)
	}
}

func main() {
	openDB()
	defer db.Close()

	pingDB()
	listTables()
}
