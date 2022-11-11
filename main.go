package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"utilities/uri"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func main() {
	router.HandleFunc("/greet/{name}", func(w http.ResponseWriter, r *http.Request) {
		var vars = mux.Vars(r)
		name := vars["name"]
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(name)
	}).Methods("GET")

	uri.Uri_init_inverter(router)
	uri.Uri_init_calculator(router)
	fmt.Println("Server started at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
