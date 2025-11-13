package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Learning json")

	type Person struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		IsAdult bool   `json:"is_adult"`
	}

	person := Person{
		Name:    "John Doe",
		Age:     30,
		IsAdult: true,
	}

	fmt.Println("Person Data : ", person)

	// convert to json
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println("JSON Data : ", string(jsonData))


	// deconding unmarshalling

	var decodedPerson Person
	err = json.Unmarshal(jsonData, &decodedPerson)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println("Decoded Person Data : ", decodedPerson)
}
