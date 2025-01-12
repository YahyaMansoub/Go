package Handlers

import (
	"FinalProject/functions"
	"FinalProject/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var authorMutex sync.Mutex

func HandleCreateAuthor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	var newAuthor model.Author
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newAuthor)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	authorMutex.Lock()
	defer authorMutex.Unlock()

	authors, err := functions.ReadAuthorsFromFile()
	if err != nil {
		http.Error(w, "Could not read from file", http.StatusInternalServerError)
		return
	}

	for _, author := range authors {
		if author.ID == newAuthor.ID {
			http.Error(w, "Author with this ID already exists", http.StatusBadRequest)
			return
		}
	}

	authors = append(authors, newAuthor)

	err = functions.WriteAuthorsToFile(authors)
	if err != nil {
		http.Error(w, "Could not write to file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "success", "message": "Author created successfully"}
	json.NewEncoder(w).Encode(response)
}

func HandleGetAuthorByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/get/author/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	authors, err := functions.ReadAuthorsFromFile()
	if err != nil {
		http.Error(w, "Could not read from file", http.StatusInternalServerError)
		return
	}

	for _, author := range authors {
		if author.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(author)
			return
		}
	}

	http.Error(w, fmt.Sprintf("Author with ID %d not found", id), http.StatusNotFound)
}

func HandleUpdateAuthor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/update/author/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updatedAuthor model.Author
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updatedAuthor)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	authorMutex.Lock()
	defer authorMutex.Unlock()

	authors, err := functions.ReadAuthorsFromFile()
	if err != nil {
		http.Error(w, "Could not read from file", http.StatusInternalServerError)
		return
	}

	var authorFound bool
	for i, author := range authors {
		if author.ID == id {
			authors[i] = updatedAuthor
			authorFound = true
			break
		}
	}

	if !authorFound {
		http.Error(w, fmt.Sprintf("Author with ID %d not found", id), http.StatusNotFound)
		return
	}

	err = functions.WriteAuthorsToFile(authors)
	if err != nil {
		http.Error(w, "Could not write to file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "success", "message": fmt.Sprintf("Author with ID %d updated successfully", id)}
	json.NewEncoder(w).Encode(response)
}

func HandleDeleteAuthor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/delete/author/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	authorMutex.Lock()
	defer authorMutex.Unlock()

	err = functions.DeleteAuthorByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "success", "message": fmt.Sprintf("Author with ID %d deleted successfully", id)}
	json.NewEncoder(w).Encode(response)
}
func HandleListAllAuthors(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	authors, err := functions.ReadAuthorsFromFile()
	if err != nil {
		http.Error(w, "Could not read from file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}
