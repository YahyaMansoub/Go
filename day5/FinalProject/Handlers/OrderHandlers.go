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

var orderMutex sync.Mutex

func HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	var newOrder model.Order
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newOrder)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	orderMutex.Lock()
	defer orderMutex.Unlock()

	orders, err := functions.ReadOrdersFromFile()
	if err != nil {
		http.Error(w, "Could not read from file", http.StatusInternalServerError)
		return
	}

	orders = append(orders, newOrder)

	err = functions.WriteOrdersToFile(orders)
	if err != nil {
		http.Error(w, "Could not write to file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "success", "message": "Order created successfully"}
	json.NewEncoder(w).Encode(response)
}

func HandleGetOrderByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/get/order/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	orders, err := functions.ReadOrdersFromFile()
	if err != nil {
		http.Error(w, "Could not read from file", http.StatusInternalServerError)
		return
	}

	for _, order := range orders {
		if order.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(order)
			return
		}
	}

	http.Error(w, fmt.Sprintf("Order with ID %d not found", id), http.StatusNotFound)
}

func HandleUpdateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/update/order/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updatedOrder model.Order
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updatedOrder)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	orderMutex.Lock()
	defer orderMutex.Unlock()

	orders, err := functions.ReadOrdersFromFile()
	if err != nil {
		http.Error(w, "Could not read from file", http.StatusInternalServerError)
		return
	}

	var orderFound bool
	for i, order := range orders {
		if order.ID == id {
			orders[i] = updatedOrder
			orderFound = true
			break
		}
	}

	if !orderFound {
		http.Error(w, fmt.Sprintf("Order with ID %d not found", id), http.StatusNotFound)
		return
	}

	err = functions.WriteOrdersToFile(orders)
	if err != nil {
		http.Error(w, "Could not write to file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "success", "message": fmt.Sprintf("Order with ID %d updated successfully", id)}
	json.NewEncoder(w).Encode(response)
}

func HandleDeleteOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/api/delete/order/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	orderMutex.Lock()
	defer orderMutex.Unlock()

	err = functions.DeleteOrderByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "success", "message": fmt.Sprintf("Order with ID %d deleted successfully", id)}
	json.NewEncoder(w).Encode(response)
}

func HandleListAllOrders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	orders, err := functions.ReadOrdersFromFile()
	if err != nil {
		http.Error(w, "Could not read from file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
