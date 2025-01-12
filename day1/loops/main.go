package main

import "fmt"

func main() {
	var sum int = 1
	for sum < 100 {
		fmt.Println("Sum is still under budget")
		sum *= 10
	}

	numbers := []int{1, 2, 3}
	for index, value := range numbers {
		fmt.Println(index, value)
	}
}
