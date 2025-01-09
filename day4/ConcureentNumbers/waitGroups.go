package main

import (
	"fmt"
	"sync"
)

func GenSlic() []int {
	var slice []int
	for i := 0; i < 10; i++ {
		slice = append(slice, i+1)
	}
	return slice
}

func squareNumber(number int, wg *sync.WaitGroup) 
	defer wg.Done() 
	result := number * number
	fmt.Println("Square of", number, "is", result)
}

func main() {
	slice := GenSlic()
	var wg sync.WaitGroup 

	for _, value := range slice {
		wg.Add(1) 
		go squareNumber(value, &wg) 
	}

	wg.Wait() 
}
