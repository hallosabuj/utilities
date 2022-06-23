package inverter

import (
	"fmt"
	"strings"
)

func CreateInverterTable() {
	db := ConnectDB()
	createTable := "CREATE TABLE APPLIANCES(id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, name VARCHAR(20), watt INT);"
	_, err := db.Exec(createTable)
	if strings.Contains((err.Error()), "1050") {
		fmt.Println("Table Exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("Table created successfully")
	}
	db.Close()
}

func DBInsertAppliances(appliances []Appliance) []string {
	var output []string
	db := ConnectDB()
	for _, appliance := range appliances {
		sqlCmd := fmt.Sprintf("INSERT INTO APPLIANCES(name,watt) values('%s', %d)", appliance.Name, appliance.Watt)
		_, err := db.Query(sqlCmd)
		if strings.Contains(err.Error(), "1062") {
			output = append(output, fmt.Sprintf("%s Already present", appliance.Name))
		} else if err != nil {
			output = append(output, fmt.Sprintf("%s Unable to Insert", appliance.Name))
		} else {
			output = append(output, fmt.Sprintf("%s Inserted successfully", appliance.Name))
		}
	}
	db.Close()
	return output
}

func DBFetchAppliances() []Appliance {
	db := ConnectDB()
	result, err := db.Query("SELECT name,watt from APPLIANCES order by watt")
	if err != nil {
		panic(err)
	}
	var appliances []Appliance
	for result.Next() {
		var name string
		var watt int
		err := result.Scan(&name, &watt)
		if err != nil {
			panic(err)
		}
		appliances = append(appliances, Appliance{Name: name, Watt: watt})
	}
	db.Close()
	return appliances
}
