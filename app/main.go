package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maybemanolo/api/search"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/pi/{number}", pihandler)
	fmt.Println(":8080 serving")
	http.ListenAndServe(":8080", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}

type out struct {
	Index int    `json:"index"`
	Text  string `json:"text"`
}

func pihandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	num := vars["number"]
	res, index := search.Get(num)
	out := out{index, res}
	json.NewEncoder(w).Encode(out)
}
