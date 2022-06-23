package uri

import (
	"net/http"
	"utilities/calculator"

	"github.com/gorilla/mux"
)

func Uri_init_calculator(router *mux.Router) {
	router.HandleFunc("/calculator/add", func(w http.ResponseWriter, r *http.Request) {
		calculator.Addition(w, r)
	}).Methods("POST")
	router.HandleFunc("/calculator/sub", func(w http.ResponseWriter, r *http.Request) {
		calculator.Subtraction(w, r)
	}).Methods("POST")
	router.HandleFunc("/calculator/mult", func(w http.ResponseWriter, r *http.Request) {
		calculator.Multiplication(w, r)
	}).Methods("POST")
	router.HandleFunc("/calculator/div", func(w http.ResponseWriter, r *http.Request) {
		calculator.Division(w, r)
	}).Methods("POST")
}
