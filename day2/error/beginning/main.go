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


func formatDivisionMessage(a,b int)(string,error){
	divisionResult , err := divide(a,b)
	if err != nil {
		return "",errors.
	}
	
}

func main(){
	var  a,b := 10,20
	defer fmt.Println("A=",a)
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