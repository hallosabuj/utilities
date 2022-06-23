package calculator

import (
	"encoding/json"
	"net/http"
)

func Multiplication(w http.ResponseWriter, r *http.Request) {
	var numbers MultNumbers
	json.NewDecoder(r.Body).Decode(&numbers)

	var result MultResult
	result.Numbers = numbers.Numbers
	result.Product = 1
	for _, num := range numbers.Numbers {
		result.Product *= num
	}
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(result)
}
