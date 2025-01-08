package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Define the User struct that will be encoded into JSON
type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// Create some sample users
	users := []User{
		{Name: "Alice", Age: 30, Email: "alice@example.com"},
		{Name: "Bob", Age: 25, Email: "bob@example.com"},
	}

	// Open a file for writing the JSON data
	file, err := os.Create("users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Ensure the file is closed after writing

	// Create a new JSON encoder
	encoder := json.NewEncoder(file)

	// Set the indentation to make the JSON pretty-printed
	encoder.SetIndent("", "  ")

	// Encode the users slice into the file
	err = encoder.Encode(users)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("JSON data has been written to users.json")
}
