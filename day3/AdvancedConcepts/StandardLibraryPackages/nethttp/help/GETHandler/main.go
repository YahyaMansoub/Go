package main

import (
	"fmt"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	// Handle GET request
	if r.Method == http.MethodGet {
		// Send a simple response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"message": "Hello, this is a GET response!"}`)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/api/get", getHandler)
	
	// Start the server
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
