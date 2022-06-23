package calculator

import (
	"encoding/json"
	"net/http"
)

func Subtraction(w http.ResponseWriter, r *http.Request) {
	var numbers SubNumbers
	json.NewDecoder(r.Body).Decode(&numbers)

	var result SubResult
	result.Number1 = numbers.Number1
	result.Number2 = numbers.Number2
	result.Difference = numbers.Number1 - result.Number2

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(result)
}
