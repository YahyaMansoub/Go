package main

import (
	"fmt"
)

func riskyFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()
	panic("Something went wrong")
}

func main() {
	riskyFunction()
	fmt.Println("Program continues...")
}
