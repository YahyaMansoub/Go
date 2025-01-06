package main

import "fmt"

func main(){
	var length float64 = 0
	var width float64 = 0


	if length == 0 {
		fmt.Println("length empty")
	}else if width == 0 {
		fmt.Println("width empty")
	}else if isLengthEmpty:= length == 0; isLengthEmpty{  //the variable is valable only within the if statement , undefined out fo the if statement 
		fmt.Println("length empty")  

	}

	grade := "A"
	switch grade {

	case "A":
		fmt.Println("You got an A!")
		fallthrough   //fallthrough : makes it possible to excecute the conditions below 
	// in Go ,  go automaticaally breaks 
	case "B":
		fmt.Println("You got an A!")
	
    default:
		fmt.Println("You got an A!")
	}

	number :=9


	switch {
	case number < 10:
		fmt.Println("less than 10")
		fallthrough
	case number < 20:
		fmt.Println("less than 20")
		fallthrough
	case number < 30:
		fmt.Println("less than 30")
		fallthrough
	default:
		fmt.Println("default case")
	}




}