package gofood

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	dbUrl := os.Getenv("DBURL")
	var repo FoodRepo

	if dbUrl == "" {
		//panic("Please set the DBURL env variable: postgres://user:pwd@host/dbname?sslmode=disable")
		repo = NewDummyFoodRepo()
	} else {
		repo = NewFoodDB(dbUrl)
	}

	fw := FoodWeb{
		Repo: repo,
	}

	http.HandleFunc("/food", fw.foodListHandler)
	http.HandleFunc("/info", infoHandler)
	http.ListenAndServe(":8080", nil)
}
