package functions

import (
	"FinalProject/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var orderMutex sync.RWMutex

func GetOrderDatabaseFilePath() (string, error) {
	absPath, err := filepath.Abs("orderdb.json")
	if err != nil {
		return "", err
	}
	return absPath, nil
}

func ReadOrdersFromFile() ([]model.Order, error) {
	orderMutex.RLock()
	defer orderMutex.RUnlock()

	filePath, err := GetOrderDatabaseFilePath()
	if err != nil {
		return nil, err
	}

	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Order{}, nil
		}
		return nil, err
	}

	var orders []model.Order
	err = json.Unmarshal(fileContent, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func WriteOrdersToFile(orders []model.Order) error {
	orderMutex.Lock()
	defer orderMutex.Unlock()

	filePath, err := GetOrderDatabaseFilePath()
	if err != nil {
		return err
	}

	fileContent, err := json.MarshalIndent(orders, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, fileContent, 0644)
}

func DeleteOrderByID(id int) error {
	orderMutex.Lock()
	defer orderMutex.Unlock()

	orders, err := ReadOrdersFromFile()
	if err != nil {
		return err
	}

	var indexToDelete = -1
	for i, order := range orders {
		if order.ID == id {
			indexToDelete = i
			break
		}
	}

	if indexToDelete == -1 {
		return fmt.Errorf("order with ID %d not found", id)
	}

	orders = append(orders[:indexToDelete], orders[indexToDelete+1:]...)
	return WriteOrdersToFile(orders)
}
