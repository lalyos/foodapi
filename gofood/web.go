package gofood

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gorilla/mux"
)

type FoodWeb struct {
	Repo FoodRepo
}

func (fw FoodWeb) foodListHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("[web] foodListHadler")
	b, _ := json.Marshal(fw.Repo.GetAllFoodList())
	w.Write(b)
}

func (fw FoodWeb) addFoodHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("[web] addFoodHandler")

	decoder := json.NewDecoder(req.Body)
	var f Food
	err := decoder.Decode(&f)
	if err != nil {
		log.Println("[WARNING] couldnt add food:", err)
	}
	fw.Repo.AddFood(f)
	fmt.Fprintf(w, "OK %v", f)
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

	r := mux.NewRouter()
	r.HandleFunc("/food", fw.foodListHandler).Methods("GET")
	r.HandleFunc("/food", fw.addFoodHandler).Methods("POST")
	r.HandleFunc("/info", infoHandler)
	log.Println("[web] starting server at:", address)
	http.Handle("/", r)
	http.ListenAndServe(address, nil)
}
