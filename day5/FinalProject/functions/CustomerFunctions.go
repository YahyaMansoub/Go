package functions

import (
	"FinalProject/model"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

var customerMutex sync.RWMutex

func ReadCustomersFromFile() ([]model.Customer, error) {
	customerMutex.RLock()
	defer customerMutex.RUnlock()

	var customers []model.Customer
	file, err := os.Open("CustomerDB.json")
	if err != nil {
		return customers, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&customers)
	if err != nil {
		return customers, err
	}

	return customers, nil
}

func WriteCustomersToFile(customers []model.Customer) error {
	customerMutex.Lock()
	defer customerMutex.Unlock()

	file, err := os.Create("CustomerDB.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(customers)
}

func DeleteCustomerByID(id int) error {
	customerMutex.Lock()
	defer customerMutex.Unlock()

	customers, err := ReadCustomersFromFile()
	if err != nil {
		return err
	}

	var indexToDelete = -1
	for i, customer := range customers {
		if customer.ID == id {
			indexToDelete = i
			break
		}
	}

	if indexToDelete == -1 {
		return fmt.Errorf("Customer with ID %d not found", id)
	}

	customers = append(customers[:indexToDelete], customers[indexToDelete+1:]...)
	return WriteCustomersToFile(customers)
}
