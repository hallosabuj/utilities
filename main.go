package main

import (
	"log"
	"net/http"
	"utilities/uri"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func main() {
	uri.Uri_init_inverter(router)
	uri.Uri_init_calculator(router)
	log.Fatal(http.ListenAndServe(":5000", router))
	// db.ConnectDB()
	// db.CreateInverterTable()
}
