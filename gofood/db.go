package gofood

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	_ "github.com/lib/pq"
)

var err error
var hostname string

type FoodDB struct {
	dbUrl string
	db    *sql.DB
}

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

func (f FoodDB) pingDB() {
	err = f.db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func removeAuthInfo(dburl string) string {
	u, _ := url.Parse(dburl)
	return u.Scheme + "://" + u.Host + u.Path + "?" + u.RawQuery
}

func (f *FoodDB) openDB() {
	log.Println("[INFO] connecting to", removeAuthInfo(f.dbUrl))
	f.db, err = sql.Open("postgres", f.dbUrl)
	if err != nil {
		log.Fatal(err)
	}
}

func (f FoodDB) getTables(schema string) (map[string]bool, error) {
	tables := map[string]bool{}

	rows, err := f.db.Query(listTablesSql, schema)
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

func (f FoodDB) listTables() {
	tables, err := f.getTables("public")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tables:", tables)
}

func (f FoodDB) insertTestData() {
	testData := map[string]int{
		"pacal":    550,
		"pancake":  400,
		"tortilla": 1400,
		"pizza":    1200,
	}

	insertStmt, err := f.db.Prepare(inserSql)
	if err != nil {
		log.Fatal(err)
	}

	for food, price := range testData {
		insertStmt.Exec(food, price)
	}
}

func (f FoodDB) Add(food Food) {
	log.Println("TODO")
}

func (f FoodDB) Delete(name string) {
	log.Println("TODO")
}

func (f FoodDB) Get(name string) (Food, bool) {
	log.Println("TODO")
	return Food{}, false
}

func (f FoodDB) Update(food Food) bool {
	log.Println("TODO")
	return false
}

func (f FoodDB) GetAllFoodList() []Food {
	foodList := []Food{}
	for _, f := range f.GetAllFoodMap() {
		foodList = append(foodList, f)
	}
	return foodList
}

func (f FoodDB) GetAllFoodMap() map[string]Food {
	foodMap := map[string]Food{}

	rows, err := f.db.Query(listFoodSql)
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

func (f FoodDB) createFoodTableIfNotExists() {
	schema, table := "public", "food"
	tables, err := f.getTables(schema)
	if err != nil {
		log.Fatal(err)
	}

	if _, exists := tables[table]; exists {
		log.Println("table already exists:", table)

	} else {
		log.Println("missing table:", table)
		_, err := f.db.Exec(createTableSql)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("create table SUCCESS:", table)
		log.Println("inserting test data ...")
		f.insertTestData()
	}
}

func init() {
	log.SetFlags(log.Ltime)
	log.SetPrefix("[INFO] ")
	hostname, _ = os.Hostname()
}

func NewFoodDB(url string) *FoodDB {
	f := FoodDB{
		dbUrl: url,
	}
	f.openDB()
	f.pingDB()
	f.createFoodTableIfNotExists()

	return &f
}
