package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	// Create a sample struct
	person := Person{
		Name:    "Alice",
		Age:     30,
		Address: "123 Main St",
	}

	// Encoding the struct into JSON format
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}

	// Print the encoded JSON data as a string
	fmt.Println(string(jsonData))
}
