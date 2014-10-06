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

func (fw FoodWeb) getFoodHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("[web] getFoodHandler")
	vars := mux.Vars(req)
	foodName := vars["food"]

	f, ok := fw.Repo.GetFood(foodName)
	var b []byte
	if ok {
		b, _ = json.Marshal(f)
	} else {
		b = []byte(`{"message": "NotFound"}`)
	}
	w.Write(b)
}

func (fw FoodWeb) deleteFoodHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("[web] deleteFoodHandler")
	vars := mux.Vars(req)
	foodName := vars["food"]

	fw.Repo.DeleteFood(foodName)
	fmt.Fprintf(w, "DELETED %v", foodName)
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

func (fw FoodWeb) updateFoodHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("[web] updateFoodHandler")

	decoder := json.NewDecoder(req.Body)
	var f Food
	err := decoder.Decode(&f)
	if err != nil {
		log.Println("[WARNING] couldnt add food:", err)
	}

	ok := fw.Repo.UpdateFood(f)
	if ok {
		fmt.Fprintf(w, "UPDATED %v", f.Name)
	} else {
		fmt.Fprintf(w, "Not Found ")
	}

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
	r.HandleFunc("/food/{food}", fw.getFoodHandler).Methods("GET")
	r.HandleFunc("/food/{food}", fw.deleteFoodHandler).Methods("DELETE")
	r.HandleFunc("/food/{food}", fw.updateFoodHandler).Methods("PUT")
	r.HandleFunc("/food", fw.addFoodHandler).Methods("POST")
	r.HandleFunc("/info", infoHandler)
	log.Println("[web] starting server at:", address)
	http.Handle("/", r)
	http.ListenAndServe(address, nil)
}
