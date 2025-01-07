package main 

import (
	"fmt"
	"errors"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}


func main(){
	var a, b int
	fmt.Println("Enter two numbers:")
	fmt.Scanln(&a, &b)
	quotient, err := divide(a, b)
	if err != nil {
		fmt.Println(err)
		} else {
			fmt.Println("Quotient:", quotient)
			}

}