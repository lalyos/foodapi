package main

import (
	"fmt"
	"os"

	"github.com/lalyos/gofood/gofood"
)

func main() {
	dbUrl := os.Getenv("DBURL")
	repo := gofood.NewFoodDB(dbUrl)
	fmt.Println(repo.GetAllFoodList())
}
