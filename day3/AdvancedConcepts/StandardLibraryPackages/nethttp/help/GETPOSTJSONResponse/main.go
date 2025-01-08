package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct for response data
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	// Handle GET request
	if r.Method == http.MethodGet {
		// Create response
		response := Response{
			Status:  "success",
			Message: "This is a GET response with JSON",
		}

		// Encode and send response as JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// Handle POST request
	if r.Method == http.MethodPost {
		var msg Response
		// Decode incoming JSON request
		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Create and send response
		msg.Status = "received"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(msg)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/api/get", getHandler)
	http.HandleFunc("/api/post", postHandler)

	// Start the server
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
