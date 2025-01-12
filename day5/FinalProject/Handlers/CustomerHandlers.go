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

var customerMutex sync.Mutex

func HandleCreateCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	var newCustomer model.Customer
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newCustomer)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	customerMutex.Lock()
	defer customerMutex.Unlock()

	customers, err := functions.ReadCustomersFromFile()
	if err != nil {
		http.Error(w, "Could not read from file", http.StatusInternalServerError)
		return
	}

	for _, customer := range customers {
		if customer.ID == newCustomer.ID {
			http.Error(w, "Customer with this ID already exists", http.StatusBadRequest)
			return
		}
	}

	customers = append(customers, newCustomer)

	err = functions.WriteCustomersToFile(customers)
	if err != nil {
		http.Error(w, "Could not write to file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "success", "message": "Customer created successfully"}
	json.NewEncoder(w).Encode(response)
}

func HandleGetCustomerByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/customers/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	customers, err := functions.ReadCustomersFromFile()
	if err != nil {
		http.Error(w, "Could not read from file", http.StatusInternalServerError)
		return
	}

	for _, customer := range customers {
		if customer.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(customer)
			return
		}
	}

	http.Error(w, fmt.Sprintf("Customer with ID %d not found", id), http.StatusNotFound)
}

func HandleUpdateCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/customers/update/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updatedCustomer model.Customer
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updatedCustomer)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	customerMutex.Lock()
	defer customerMutex.Unlock()

	customers, err := functions.ReadCustomersFromFile()
	if err != nil {
		http.Error(w, "Could not read from file", http.StatusInternalServerError)
		return
	}

	var customerFound bool
	for i, customer := range customers {
		if customer.ID == id {
			customers[i] = updatedCustomer
			customerFound = true
			break
		}
	}

	if !customerFound {
		http.Error(w, fmt.Sprintf("Customer with ID %d not found", id), http.StatusNotFound)
		return
	}

	err = functions.WriteCustomersToFile(customers)
	if err != nil {
		http.Error(w, "Could not write to file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "success", "message": fmt.Sprintf("Customer with ID %d updated successfully", id)}
	json.NewEncoder(w).Encode(response)
}

func HandleDeleteCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/customers/delete/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	customerMutex.Lock()
	defer customerMutex.Unlock()

	err = functions.DeleteCustomerByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "success", "message": fmt.Sprintf("Customer with ID %d deleted successfully", id)}
	json.NewEncoder(w).Encode(response)
}

func HandleListAllCustomers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	customers, err := functions.ReadCustomersFromFile()
	if err != nil {
		http.Error(w, "Could not read from file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
