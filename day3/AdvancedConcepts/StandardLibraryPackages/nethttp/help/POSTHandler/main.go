package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Define a struct to handle POST body
type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// Handle POST request
	if r.Method == http.MethodPost {
		var msg Message
		// Decode the incoming JSON request
		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Respond with a success message
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := fmt.Sprintf(`{"status": "received", "name": "%s", "message": "%s"}`, msg.Name, msg.Message)
		fmt.Fprintf(w, response)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/api/post", postHandler)
	
	// Start the server
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
