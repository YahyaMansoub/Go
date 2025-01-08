package main

import (
	"fmt"
	"errors"
)

func readFile(filename string) error {
	return fmt.Errorf("reading file %s: %w", filename, errors.New("file not found"))
}

func main() {
	err := readFile("example.txt")
	if err != nil {
		// Unwrap the error to access the original error
		fmt.Println("Error:", err)
	}
}
