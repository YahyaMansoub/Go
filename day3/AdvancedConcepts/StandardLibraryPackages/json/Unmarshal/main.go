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
	// JSON data to be decoded
	jsonData := `{"name":"Bob","age":25,"address":"456 Elm St"}`

	// Declare a variable to hold the decoded data
	var person Person

	// Decoding the JSON data into the 'person' variable
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		log.Fatal(err)
	}

	// Print the decoded struct
	fmt.Printf("%+v\n", person)
}
