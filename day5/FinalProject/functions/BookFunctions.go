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

var bookMutex sync.RWMutex

func GetDatabaseFilePath() (string, error) {
	absPath, err := filepath.Abs("database.json")
	if err != nil {
		return "", err
	}
	return absPath, nil
}

func ReadBooksFromFile() ([]model.Book, error) {
	bookMutex.RLock()
	defer bookMutex.RUnlock()

	filePath, err := GetDatabaseFilePath()
	if err != nil {
		return nil, err
	}

	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Book{}, nil
		}
		return nil, err
	}

	var books []model.Book
	err = json.Unmarshal(fileContent, &books)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func WriteBooksToFile(books []model.Book) error {
	bookMutex.Lock()
	defer bookMutex.Unlock()

	filePath, err := GetDatabaseFilePath()
	if err != nil {
		return err
	}

	fileContent, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, fileContent, 0644)
}

func DeleteBookByID(id int) error {
	bookMutex.Lock()
	defer bookMutex.Unlock()

	books, err := ReadBooksFromFile()
	if err != nil {
		return err
	}

	var indexToDelete = -1
	for i, book := range books {
		if book.ID == id {
			indexToDelete = i
			break
		}
	}

	if indexToDelete == -1 {
		return fmt.Errorf("book with ID %d not found", id)
	}

	books = append(books[:indexToDelete], books[indexToDelete+1:]...)
	return WriteBooksToFile(books)
}
