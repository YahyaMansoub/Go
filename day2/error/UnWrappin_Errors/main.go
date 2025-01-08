package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func search() error {
	return fmt.Errorf("search failed: %w", ErrNotFound)
}

func main() {
	err := search()
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Custom error occurred:", err)
	}
}
