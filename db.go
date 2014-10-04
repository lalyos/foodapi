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

const listTablesSql = `
  SELECT table_name
    FROM information_schema.tables
    WHERE table_schema=$1
    ORDER BY table_name;
`

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

func getTables(schema string) (map[string]bool, error) {
	tables := map[string]bool{}

	rows, err := db.Query(listTablesSql, schema)
	if err != nil {
		return tables, err
	}
	var table string
	for rows.Next() {
		err := rows.Scan(&table)
		if err != nil {
			log.Fatal(err)
		}
		tables[table] = true
	}
	return tables, nil
}

func listTables() {
	tables, err := getTables("public")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tables:", tables)
}

func main() {
	openDB()
	defer db.Close()

func init() {
	log.SetFlags(log.Ltime)
	log.SetPrefix("[INFO] ")
}

	pingDB()
	listTables()
}
