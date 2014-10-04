package gofood

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error
var hostname string

const version = "1.0"
const createTableSql = `
  CREATE TABLE food (
    name    varchar(40) NOT NULL,
    price       integer NOT NULL
  );
`

const listTablesSql = `
  SELECT table_name
    FROM information_schema.tables
    WHERE table_schema=$1
    ORDER BY table_name;
`

const inserSql = "INSERT into food VALUES ($1, $2)"
const listFoodSql = "SELECT * from food;"

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
	defer rows.Close()
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

func insertTestData() {
	testData := map[string]int{
		"pacal":    550,
		"pancake":  400,
		"tortilla": 1400,
		"pizza":    1200,
	}

	insertStmt, err := db.Prepare(inserSql)
	if err != nil {
		log.Fatal(err)
	}

	for food, price := range testData {
		insertStmt.Exec(food, price)
	}
}

func GetAllFoodList() []Food {
	foodList := []Food{}
	for _, f := range GetAllFoodMap() {
		foodList = append(foodList, f)
	}
	return foodList
}

func GetAllFoodMap() map[string]Food {
	foodMap := map[string]Food{}

	rows, err := db.Query(listFoodSql)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		f := Food{}
		err := rows.Scan(&f.Name, &f.Price)
		if err != nil {
			log.Fatal(err)
		}
		foodMap[f.Name] = f
	}

	return foodMap
}

func createFoodTableIfNotExists() {
	schema, table := "public", "food"
	tables, err := getTables(schema)
	if err != nil {
		log.Fatal(err)
	}

	if _, exists := tables[table]; exists {
		log.Println("table already exists:", table)

	} else {
		log.Println("missing table:", table)
		_, err := db.Exec(createTableSql)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("create table SUCCESS:", table)
		log.Println("inserting test data ...")
		insertTestData()
	}
}

func init() {
	log.SetFlags(log.Ltime)
	log.SetPrefix("[INFO] ")
	hostname, _ = os.Hostname()
}

func InitDb() {
	openDB()
	pingDB()
	createFoodTableIfNotExists()
}
