package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Define the User struct that matches the JSON structure
type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// Open the JSON file
	file, err := os.Open("users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Ensure the file is closed after reading

	// Declare a slice to hold the decoded users
	var users []User

	// Create a JSON decoder
	decoder := json.NewDecoder(file)

	// Decode the JSON content into the users slice
	err = decoder.Decode(&users)
	if err != nil {
		log.Fatal(err)
	}

	// Print all the users
	fmt.Println("Decoded Users:")
	for _, user := range users {
		fmt.Printf("Name: %s, Age: %d, Email: %s\n", user.Name, user.Age, user.Email)
	}

	// Extract a specific user (e.g., find user by name)
	for _, user := range users {
		if user.Name == "Alice" {
			fmt.Printf("\nFound Alice: %+v\n", user)
		}
	}
}
