package uri

import (
	"net/http"
	"utilities/inverter"

	"github.com/gorilla/mux"
)

func Uri_init_inverter(router *mux.Router) {
	router.HandleFunc("/inverter/add-appliances", func(w http.ResponseWriter, r *http.Request) {
		inverter.AddAppliances(w, r)
	}).Methods("POST")

	router.HandleFunc("/inverter/fetch-appliances", func(w http.ResponseWriter, r *http.Request) {
		inverter.FetchAppliances(w, r)
	}).Methods("GET")

	router.HandleFunc("/inverter/calculate-details", func(w http.ResponseWriter, r *http.Request) {
		inverter.CalculateDetails(w, r)
	}).Methods("POST")
}
