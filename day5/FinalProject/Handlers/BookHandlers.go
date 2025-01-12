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

var mutex sync.Mutex

func HandleCreateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	var newBook model.Book
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newBook)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	books, err := functions.ReadBooksFromFile()
	if err != nil {
		http.Error(w, "Could not read from file", http.StatusInternalServerError)
		return
	}

	for _, book := range books {
		if book.ID == newBook.ID {
			http.Error(w, "Book with this ID already exists", http.StatusBadRequest)
			return
		}
	}

	books = append(books, newBook)

	err = functions.WriteBooksToFile(books)
	if err != nil {
		http.Error(w, "Could not write to file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "success", "message": "Book created successfully"}
	json.NewEncoder(w).Encode(response)
}

func HandleGetBookByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/get/book/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	books, err := functions.ReadBooksFromFile()
	if err != nil {
		http.Error(w, "Could not read from file", http.StatusInternalServerError)
		return
	}

	for _, book := range books {
		if book.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	http.Error(w, fmt.Sprintf("Book with ID %d not found", id), http.StatusNotFound)
}

func HandleUpdateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/update/book/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updatedBook model.Book
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	books, err := functions.ReadBooksFromFile()
	if err != nil {
		http.Error(w, "Could not read from file", http.StatusInternalServerError)
		return
	}

	var bookFound bool
	for i, book := range books {
		if book.ID == id {
			books[i] = updatedBook
			bookFound = true
			break
		}
	}

	if !bookFound {
		http.Error(w, fmt.Sprintf("Book with ID %d not found", id), http.StatusNotFound)
		return
	}

	err = functions.WriteBooksToFile(books)
	if err != nil {
		http.Error(w, "Could not write to file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "success", "message": fmt.Sprintf("Book with ID %d updated successfully", id)}
	json.NewEncoder(w).Encode(response)
}

func HandleDeleteBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/delete/book/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	err = functions.DeleteBookByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "success", "message": fmt.Sprintf("Book with ID %d deleted successfully", id)}
	json.NewEncoder(w).Encode(response)
}

func HandleListAllBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	books, err := functions.ReadBooksFromFile()
	if err != nil {
		http.Error(w, "Could not read from file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func HandleSearchBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query().Get("query")
	books, err := functions.ReadBooksFromFile()
	if err != nil {
		http.Error(w, "Could not read from file", http.StatusInternalServerError)
		return
	}

	var result []model.Book
	for _, book := range books {
		if query == "" || contains(book.Title, query) || contains(book.Author.FirstName, query) || contains(book.Author.LastName, query) {
			result = append(result, book)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && (len(s) >= len(substr)) && (s[:len(substr)] == substr)
}
