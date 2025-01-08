package main

import (
	"errors"
	"fmt"
)

func validateInput(input string) error {
	if input == "" {
		return errors.New("input cannot be empty")
	}
	return nil
}

func saveToDatabase(input string) error {
	if input == "bad" {
		return errors.New("database error: bad data")
	}
	return nil
}

func process(input string) error {
	if err := validateInput(input); err != nil {
		return fmt.Errorf("validateInput failed: %w", err)
	}
	if err := saveToDatabase(input); err != nil {
		return fmt.Errorf("saveToDatabase failed: %w", err)
	}
	return nil
}

func main() {
	input := "bad"
	err := process(input)
	if err != nil {
		fmt.Println("Error occurred:", err)
	}
}
