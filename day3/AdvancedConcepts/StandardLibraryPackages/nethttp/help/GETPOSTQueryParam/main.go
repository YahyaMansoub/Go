package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	// Handle GET request with query parameters
	if r.Method == http.MethodGet {
		// Retrieve query parameters
		name := r.URL.Query().Get("name")
		message := r.URL.Query().Get("message")

		// Send response with the query parameters
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"name": "%s", "message": "%s"}`, name, message)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// Handle POST request with JSON body
	if r.Method == http.MethodPost {
		var msg map[string]string
		// Decode the incoming JSON request
		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Respond with the decoded data
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := fmt.Sprintf(`{"status": "received", "name": "%s", "message": "%s"}`, msg["name"], msg["message"])
		fmt.Fprintf(w, response)
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
