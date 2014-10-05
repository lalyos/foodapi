package gofood

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
)

type FoodWeb struct {
	Repo FoodRepo
}

func (fw FoodWeb) foodListHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("[web] foodListHadler")
	b, _ := json.Marshal(fw.Repo.GetAllFoodList())
	w.Write(b)
}

func infoHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("[web] infoHandler")
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
	address := ":8080"

	port := os.Getenv("PORT")
	if port != "" {
		address = fmt.Sprintf(":%s", port)
	}
	var repo FoodRepo

	if dbUrl == "" {
		repo = NewDummyFoodRepo()
		log.Println("[WARNING] DBURL env variable is unset, in memory repo is used")
		log.Println("[WARNING] to use postgress as backend:")
		log.Println("[WARNING]")
		log.Println("[WARNING]    export DBURL=postgres://user:pwd@host/dbname?sslmode=disable")
	} else {
		repo = NewFoodDB(dbUrl)
	}

	fw := FoodWeb{
		Repo: repo,
	}

	http.HandleFunc("/food", fw.foodListHandler)
	http.HandleFunc("/info", infoHandler)
	log.Println("[web] starting server at:", address)
	http.ListenAndServe(address, nil)
}
