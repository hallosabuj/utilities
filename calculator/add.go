package calculator

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Addition(w http.ResponseWriter, r *http.Request) {
	var numbers AddNumbers
	json.NewDecoder(r.Body).Decode(&numbers)
	fmt.Println("Numbers:", numbers)
	var result AddResult
	result.Numbers = numbers.Numbers
	result.Sum = 0
	for _, num := range numbers.Numbers {
		result.Sum += num
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(result)
}
