package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Define a struct for the request body
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the incoming JSON body
	var user User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	// Print the parsed user data
	fmt.Printf("Received: %+v\n", user)

	// Create a response
	response := map[string]string{
		"message": fmt.Sprintf("Hello, %s! We have received your email: %s", user.Username, user.Email),
	}

	// Set content type and send the response as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/submit", handlePostRequest)

	// Start the server
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
