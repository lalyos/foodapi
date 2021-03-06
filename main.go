package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/lalyos/foodapi/gofood"
)

func main() {
	if len(os.Args) > 1 && strings.HasSuffix(os.Args[1], "version") {
		fmt.Println(gofood.Version)
		return
	}
	gofood.NewDBBasedFoodWeb()
}
