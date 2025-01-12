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

var authorMutex sync.RWMutex

func GetAuthorsFilePath() (string, error) {
	absPath, err := filepath.Abs("authors.json")
	if err != nil {
		return "", err
	}
	return absPath, nil
}

func ReadAuthorsFromFile() ([]model.Author, error) {
	authorMutex.RLock()
	defer authorMutex.RUnlock()

	filePath, err := GetAuthorsFilePath()
	if err != nil {
		return nil, err
	}

	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Author{}, nil
		}
		return nil, err
	}

	var authors []model.Author
	err = json.Unmarshal(fileContent, &authors)
	if err != nil {
		return nil, err
	}

	return authors, nil
}

func WriteAuthorsToFile(authors []model.Author) error {
	authorMutex.Lock()
	defer authorMutex.Unlock()

	filePath, err := GetAuthorsFilePath()
	if err != nil {
		return err
	}

	fileContent, err := json.MarshalIndent(authors, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, fileContent, 0644)
}

func DeleteAuthorByID(id int) error {
	authorMutex.Lock()
	defer authorMutex.Unlock()

	authors, err := ReadAuthorsFromFile()
	if err != nil {
		return err
	}

	var indexToDelete = -1
	for i, author := range authors {
		if author.ID == id {
			indexToDelete = i
			break
		}
	}

	if indexToDelete == -1 {
		return fmt.Errorf("author with ID %d not found", id)
	}

	authors = append(authors[:indexToDelete], authors[indexToDelete+1:]...)
	return WriteAuthorsToFile(authors)
}
