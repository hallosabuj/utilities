package calculator

import (
	"encoding/json"
	"net/http"
)

func Division(w http.ResponseWriter, r *http.Request) {
	var numbers DivNumbers
	json.NewDecoder(r.Body).Decode(&numbers)

	var result DivResult
	result.Number1 = numbers.Number1
	result.Number2 = numbers.Number2
	result.Quotient = numbers.Number1 / numbers.Number2
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(result)
}
