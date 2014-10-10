package main

import (
	"fmt"
	"os"

	"github.com/lalyos/foodapi/gofood"
)

func main() {
	dbUrl := os.Getenv("DATABASE_URL")
	repo := gofood.NewFoodDB(dbUrl)
	fmt.Println(repo.GetAllFoodList())
}
