package gofood

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
)

type FoodWeb struct {
	Repo FoodRepo
}

func (fw FoodWeb) foodListHandler(w http.ResponseWriter, req *http.Request) {
	b, _ := json.Marshal(fw.Repo.GetAllFoodList())
	w.Write(b)
}

func infoHandler(w http.ResponseWriter, req *http.Request) {
	info := fmt.Sprintf(`{"host": "%s", "version": "%s", "os":"%s", "arch":"%s" }`,
		hostname,
		version,
		runtime.GOOS,
		runtime.GOARCH,
	)
	io.WriteString(w, info)
}

func NewDBBasedFoodWeb() {
	fw := FoodWeb{}
	fw.Repo = NewFoodDB()

	http.HandleFunc("/food", fw.foodListHandler)
	http.HandleFunc("/info", infoHandler)
	http.ListenAndServe(":8080", nil)
}
