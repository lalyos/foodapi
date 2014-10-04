package main

import (
	"fmt"

	"github.com/lalyos/gofood"
)

func main() {
	gofood.InitDb()
	fmt.Println(gofood.GetAllFoodList())
}
