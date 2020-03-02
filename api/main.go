package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maybemanolo/api/search"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/pi/{number}", pihandler)
	http.ListenAndServe(":8080", router)
}

func pihandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	num := vars["number"]
	res, index := search.Get(num)
	o := out{index, res}
	json.NewEncoder(w).Encode(o)
}

type out struct {
	Index int    `json:"index"`
	Text  string `json:"text"`
}
