package main

import (
	"fmt"

	"github.com/lalyos/gofood"
)

func _main() {
	gofood.InitDb()
	fmt.Println(gofood.GetAllFoodList())
}
