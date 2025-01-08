package main

import (
	"fmt"
)

// Custom error type
type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func someFunction() error {
	// Returning a custom error
	return &MyError{Code: 404, Message: "Resource not found"}
}

func main() {
	err := someFunction()
	if err != nil {
		fmt.Println("Custom Error:", err)
	}
}
