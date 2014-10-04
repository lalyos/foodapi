package gofood

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
)

func foodListHandler(w http.ResponseWriter, req *http.Request) {
	b, _ := json.Marshal(GetAllFoodList())
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

func init() {
	InitDb()
}

func StartWeb() {
	http.HandleFunc("/food", foodListHandler)
	http.HandleFunc("/info", infoHandler)
	http.ListenAndServe(":8080", nil)
}
