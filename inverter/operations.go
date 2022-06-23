package inverter

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CalculateTotalLoad(userNeeds []UserNeed) int {
	appliancesFromDB := DBFetchAppliances()
	appliances := make(map[string]int)
	// Converting to a map
	for _, obj := range appliancesFromDB {
		appliances[obj.Name] = obj.Watt
	}

	totalLoad := 0
	for _, userNeed := range userNeeds {
		totalLoad += appliances[userNeed.ApplianceName] * userNeed.ApplianceCount
	}
	return totalLoad
}

func CalculateVARating(totalLoad int, powerFactor float32) float32 {
	return float32(totalLoad) / powerFactor
}
func CalculateBatterySize(vaRatingOfInverter float32, backupTime int, inverterPower float32) (float32, float32) {
	totalVAh := vaRatingOfInverter * float32(backupTime)
	batterySize := totalVAh / float32(inverterPower)
	return totalVAh, batterySize
}

func AddAppliances(w http.ResponseWriter, r *http.Request) {
	log.Println("Add appliances")
	var appliances []Appliance
	json.NewDecoder(r.Body).Decode(&appliances)
	fmt.Println(appliances)
	output := DBInsertAppliances(appliances)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func FetchAppliances(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetch Appliances")
	appliances := DBFetchAppliances()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appliances)
}

func CalculateDetails(w http.ResponseWriter, r *http.Request) {
	log.Println("Calculating all details")
	var userInput UserInput
	json.NewDecoder(r.Body).Decode(&userInput)
	fmt.Println(userInput)

	var result Results
	result.BackupTime = userInput.BackupTime
	result.InverterPower = 12
	result.TotalLoad = CalculateTotalLoad(userInput.ApplianceDetails)
	result.PowerFactor = 0.8
	result.VARatingOfInverter = CalculateVARating(result.TotalLoad, result.PowerFactor)
	result.TotalVAh, result.BatterySize = CalculateBatterySize(result.VARatingOfInverter, result.BackupTime, result.InverterPower)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
