package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// person struct
type person struct {
	FirstName    string  `json:"first_name"`
	Age          int     `json:"age,omitempty"` //omitempty fields
	EmailAddress string  `json:"email"`
	Address      address `json:"address"`
}

type address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

// main function
func main() {
	P1 := person{
		FirstName: "Stan",
		Age:       23,
		Address: address{
			City:  "Ohio",
			State: "London",
		},
	}

	// P2 := person{
	// 	FirstName:    "Yellow",
	// 	Age:          25,
	// 	EmailAddress: "yahoo.com",
	// }

	p1Bytes, err := json.Marshal(P1)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(p1Bytes))

	jsonData := `{"full_name":"John Doe", "emp_uid":"009", "age":30, "address":{"city": "san jose", "state": "CA"}}`

	var employeeFromJson Employee
	err = json.Unmarshal([]byte(jsonData), &employeeFromJson)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(employeeFromJson)

	listOfCityState := []address{
		{City: "Andreas", State: "ON"},
		{City: "San Jose", State: "CA"},
		{City: "Las Vegas", State: "NV"},
	}

	jsonlistOfCityState, err := json.Marshal(listOfCityState)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(jsonlistOfCityState))

	// handling unknow json structures
	jsonData2 := `{
	"name": "John", 
	"age": 30, 
	"assest": "car", 
	"address":{"city": "kent"}
	}`

	var data map[string]interface{}

	err = json.Unmarshal([]byte(jsonData2), &data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Unknown json structure:", data)

}

type Employee struct {
	FullName   string  `json:"full_name"`
	EmployeeID string  `json:"emp_uid"`
	Age        int     `json:"age"`
	Address    address `json:"address"`
}
